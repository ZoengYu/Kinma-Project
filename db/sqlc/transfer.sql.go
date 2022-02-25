// Code generated by sqlc. DO NOT EDIT.
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_product_id,
  amount
) VALUES (
  $1, $2, $3
) RETURNING id, from_account_id, to_product_id, amount, created_at, success
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToProductID   int64 `json:"to_product_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToProductID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToProductID,
		&i.Amount,
		&i.CreatedAt,
		&i.Success,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_product_id, amount, created_at, success FROM transfers
WHERE 
from_account_id = $1 OR to_product_id = $2 
LIMIT 1
`

type GetTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToProductID   int64 `json:"to_product_id"`
}

func (q *Queries) GetTransfer(ctx context.Context, arg GetTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, arg.FromAccountID, arg.ToProductID)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToProductID,
		&i.Amount,
		&i.CreatedAt,
		&i.Success,
	)
	return i, err
}

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_id, to_product_id, amount, created_at, success FROM transfers
WHERE 
    from_account_id = $1 OR
    to_product_id = $2
ORDER BY id
LIMIT $3
OFFSET $4
`

type ListTransfersParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToProductID   int64 `json:"to_product_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers,
		arg.FromAccountID,
		arg.ToProductID,
		arg.Limit,
		arg.Offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToProductID,
			&i.Amount,
			&i.CreatedAt,
			&i.Success,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransferSuccess = `-- name: UpdateTransferSuccess :one
UPDATE transfers SET success = $2
WHERE id = $1
RETURNING id, from_account_id, to_product_id, amount, created_at, success
`

type UpdateTransferSuccessParams struct {
	ID      int64 `json:"id"`
	Success bool  `json:"success"`
}

func (q *Queries) UpdateTransferSuccess(ctx context.Context, arg UpdateTransferSuccessParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, updateTransferSuccess, arg.ID, arg.Success)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToProductID,
		&i.Amount,
		&i.CreatedAt,
		&i.Success,
	)
	return i, err
}
