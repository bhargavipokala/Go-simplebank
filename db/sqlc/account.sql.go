// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  balance,
  currency,
  status
) VALUES (
  $1, $2, $3, $4
) RETURNING id, owner, balance, currency, status, created_at
`

type CreateAccountParams struct {
	Owner    string        `json:"owner"`
	Balance  int64         `json:"balance"`
	Currency string        `json:"currency"`
	Status   AccountStatus `json:"status"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.Owner,
		arg.Balance,
		arg.Currency,
		arg.Status,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, balance, currency, status, created_at FROM accounts
WHERE id = $1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getAccountForUpdate = `-- name: GetAccountForUpdate :one
SELECT id, owner, balance, currency, status, created_at FROM accounts
WHERE id = $1
FOR NO KEY UPDATE
`

func (q *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountForUpdate, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, balance, currency, status, created_at FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
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
			&i.Balance,
			&i.Currency,
			&i.Status,
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

const updateAccountBalance = `-- name: UpdateAccountBalance :one
UPDATE accounts
SET balance = balance + $1
WHERE id = $2
RETURNING id, owner, balance, currency, status, created_at
`

type UpdateAccountBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateAccountBalance(ctx context.Context, arg UpdateAccountBalanceParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountBalance, arg.Amount, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const updateAccountStatus = `-- name: UpdateAccountStatus :one
UPDATE accounts
SET status = $2
WHERE id = $1
RETURNING id, owner, balance, currency, status, created_at
`

type UpdateAccountStatusParams struct {
	ID     int64         `json:"id"`
	Status AccountStatus `json:"status"`
}

func (q *Queries) UpdateAccountStatus(ctx context.Context, arg UpdateAccountStatusParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccountStatus, arg.ID, arg.Status)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
