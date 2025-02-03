package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Transaction interface {
	Querier
	TransferTxn(ctx context.Context, request *TransferTxnRequest) (*TransferTxnResponse, error)
}

type SqlStore struct {
	*Queries
	db *sql.DB
}

func NewTransaction(db *sql.DB) Transaction {
	return &SqlStore{
		db:      db,
		Queries: New(db),
	}
}

func (t *SqlStore) execTransaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := t.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxnRequest struct {
	FromAccountID int64
	ToAccountID   int64
	Amount        int64
}

type TransferTxnResponse struct {
	Transfer    Transfer
	FromAccount Account
	ToAccount   Account
	FromEntry   Entry
	ToEntry     Entry
}

func (t *SqlStore) TransferTxn(ctx context.Context, request *TransferTxnRequest) (*TransferTxnResponse, error) {
	var response TransferTxnResponse

	err := t.execTransaction(ctx, func(q *Queries) error {
		var err error
		response.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: request.FromAccountID,
			ToAccountID:   request.ToAccountID,
			Amount:        request.Amount,
		})
		if err != nil {
			return err
		}
		response.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: request.FromAccountID,
			Amount:    -request.Amount,
		})
		if err != nil {
			return err
		}
		response.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: request.ToAccountID,
			Amount:    request.Amount,
		})
		if err != nil {
			return err
		}
		response.FromAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
			ID:     request.FromAccountID,
			Amount: -request.Amount,
		})
		if err != nil {
			return err
		}
		response.ToAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
			ID:     request.ToAccountID,
			Amount: request.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &response, nil
}
