// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: payment.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (
    name, type, logo
) VALUES (
    $1, $2, $3
) RETURNING id, name, type, logo, created_at, updated_at
`

type CreatePaymentParams struct {
	Name string      `json:"name"`
	Type string      `json:"type"`
	Logo pgtype.Text `json:"logo"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment, arg.Name, arg.Type, arg.Logo)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Logo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = $1
`

func (q *Queries) DeletePayment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deletePayment, id)
	return err
}

const getPayment = `-- name: GetPayment :one
SELECT id, name, type, logo, created_at, updated_at FROM payments
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetPayment(ctx context.Context, id int64) (Payment, error) {
	row := q.db.QueryRow(ctx, getPayment, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Logo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listPayments = `-- name: ListPayments :many
SELECT id, name, type, logo, created_at, updated_at FROM payments
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPaymentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPayments(ctx context.Context, arg ListPaymentsParams) ([]Payment, error) {
	rows, err := q.db.Query(ctx, listPayments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Payment{}
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.Logo,
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

const updatePayment = `-- name: UpdatePayment :one
UPDATE payments
SET
    name = COALESCE($2, name),
    type = COALESCE($3, type),
    logo = COALESCE($4, logo)
WHERE id = $1
RETURNING id, name, type, logo, created_at, updated_at
`

type UpdatePaymentParams struct {
	ID   int64       `json:"id"`
	Name pgtype.Text `json:"name"`
	Type pgtype.Text `json:"type"`
	Logo pgtype.Text `json:"logo"`
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, updatePayment,
		arg.ID,
		arg.Name,
		arg.Type,
		arg.Logo,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.Logo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
