// Code generated by MockGen. DO NOT EDIT.
// Source: order.go
//
// Generated by this command:
//
//	mockgen -source=order.go -destination=mock/order.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/bagashiz/go-pos/internal/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderRepositoryMockRecorder) CreateOrder(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderRepository)(nil).CreateOrder), ctx, order)
}

// GetOrderByID mocks base method.
func (m *MockOrderRepository) GetOrderByID(ctx context.Context, id uint64) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", ctx, id)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockOrderRepositoryMockRecorder) GetOrderByID(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderByID), ctx, id)
}

// ListOrders mocks base method.
func (m *MockOrderRepository) ListOrders(ctx context.Context, skip, limit uint64) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", ctx, skip, limit)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockOrderRepositoryMockRecorder) ListOrders(ctx, skip, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockOrderRepository)(nil).ListOrders), ctx, skip, limit)
}

// MockOrderService is a mock of OrderService interface.
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService.
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance.
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrderService) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, order)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrderServiceMockRecorder) CreateOrder(ctx, order any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderService)(nil).CreateOrder), ctx, order)
}

// GetOrder mocks base method.
func (m *MockOrderService) GetOrder(ctx context.Context, id uint64) (*domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrder", ctx, id)
	ret0, _ := ret[0].(*domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrder indicates an expected call of GetOrder.
func (mr *MockOrderServiceMockRecorder) GetOrder(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrder", reflect.TypeOf((*MockOrderService)(nil).GetOrder), ctx, id)
}

// ListOrders mocks base method.
func (m *MockOrderService) ListOrders(ctx context.Context, skip, limit uint64) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListOrders", ctx, skip, limit)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrders indicates an expected call of ListOrders.
func (mr *MockOrderServiceMockRecorder) ListOrders(ctx, skip, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrders", reflect.TypeOf((*MockOrderService)(nil).ListOrders), ctx, skip, limit)
}