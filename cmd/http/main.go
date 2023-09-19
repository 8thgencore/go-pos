package main

import (
	"context"
	"log/slog"
	"os"

	cache "github.com/bagashiz/go-pos/internal/adapter/cache/redis"
	handler "github.com/bagashiz/go-pos/internal/adapter/handler/http"
	repo "github.com/bagashiz/go-pos/internal/adapter/repository/postgres"
	token "github.com/bagashiz/go-pos/internal/adapter/token/paseto"
	"github.com/bagashiz/go-pos/internal/core/service"
	"github.com/joho/godotenv"
)

func init() {
	// Init logger
	var logHandler *slog.JSONHandler

	env := os.Getenv("APP_ENV")
	if env == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, nil)
	} else {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})

		// Load .env file
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", "error", err)
			os.Exit(1)
		}
	}

	logger := slog.New(logHandler)
	slog.SetDefault(logger)
}

func main() {
	appName := os.Getenv("APP_NAME")
	env := os.Getenv("APP_ENV")
	dbConn := os.Getenv("DB_CONNECTION")
	tokenSymmetricKey := os.Getenv("TOKEN_SYMMETRIC_KEY")
	httpUrl := os.Getenv("HTTP_URL")
	httpPort := os.Getenv("HTTP_PORT")
	listenAddr := httpUrl + ":" + httpPort

	slog.Info("Starting the application", "app", appName, "env", env)

	// Init database
	ctx := context.Background()
	db, err := repo.NewDB(ctx)
	if err != nil {
		slog.Error("Error initializing database connection", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database", "db", dbConn)

	// Init cache service
	cache, err := cache.NewCache(ctx)
	if err != nil {
		slog.Error("Error initializing cache connection", "error", err)
		os.Exit(1)
	}
	defer cache.Close()

	// Init token service
	token, err := token.NewToken(tokenSymmetricKey)
	if err != nil {
		slog.Error("Error initializing token service", "error", err)
		os.Exit(1)
	}

	// Dependency injection
	// User
	userRepo := repo.NewUserRepository(db)
	userService := service.NewUserService(userRepo, cache)
	userHandler := handler.NewUserHandler(userService)

	// Auth
	authService := service.NewAuthService(userRepo, token)
	authHandler := handler.NewAuthHandler(authService)

	// Payment
	paymentRepo := repo.NewPaymentRepository(db)
	paymentService := service.NewPaymentService(paymentRepo, cache)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// Category
	categoryRepo := repo.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo, cache)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Product
	productRepo := repo.NewProductRepository(db)
	productService := service.NewProductService(productRepo, categoryRepo, cache)
	productHandler := handler.NewProductHandler(productService)

	// Order
	orderRepo := repo.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo, productRepo, categoryRepo, userRepo, paymentRepo, cache)
	orderHandler := handler.NewOrderHandler(orderService)

	// Init router
	router, err := handler.NewRouter(
		token,
		*userHandler,
		*authHandler,
		*paymentHandler,
		*categoryHandler,
		*productHandler,
		*orderHandler,
	)
	if err != nil {
		slog.Error("Error initializing router", "error", err)
		os.Exit(1)
	}

	// Start server
	slog.Info("Starting the HTTP server", "listen_address", listenAddr)
	err = router.Serve(listenAddr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
		os.Exit(1)
	}
}
