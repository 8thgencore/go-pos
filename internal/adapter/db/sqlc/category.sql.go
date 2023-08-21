// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: category.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
    name
) VALUES (
    $1
) RETURNING id, name, created_at, updated_at
`

func (q *Queries) CreateCategory(ctx context.Context, name string) (Category, error) {
	row := q.db.QueryRow(ctx, createCategory, name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, created_at, updated_at FROM categories
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRow(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, created_at, updated_at FROM categories
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.Query(ctx, listCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategory = `-- name: UpdateCategory :one
UPDATE categories
SET
    name = COALESCE($2, name)
WHERE id = $1
RETURNING id, name, created_at, updated_at
`

type UpdateCategoryParams struct {
	ID   int64       `json:"id"`
	Name pgtype.Text `json:"name"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, updateCategory, arg.ID, arg.Name)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
