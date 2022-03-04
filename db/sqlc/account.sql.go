// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO account (
  owner,
  currency
) VALUES (
  $1, $2
) RETURNING id, owner, currency, created_at
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount, arg.Owner, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM account
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, currency, created_at FROM account
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listAccount = `-- name: ListAccount :many
SELECT id, owner, currency, created_at FROM account
WHERE owner = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListAccountParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) ListAccount(ctx context.Context, arg ListAccountParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccount, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Currency,
			&i.CreatedAt,
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

const updateAccount = `-- name: UpdateAccount :one
UPDATE account SET currency =$2
WHERE id = $1
RETURNING id, owner, currency, created_at
`

type UpdateAccountParams struct {
	ID       int64  `json:"id"`
	Currency string `json:"currency"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount, arg.ID, arg.Currency)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}