package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTxn(t *testing.T) {
	transaction := NewTransaction(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxnResponse)

	for i := 0; i < n; i++ {
		go func() {
			result, err := transaction.TransferTxn(context.Background(), &TransferTxnRequest{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- *result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		transfer := <-results
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.Transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.Transfer.ToAccountID, account2.ID)
		require.Equal(t, transfer.Transfer.Amount, amount)
		require.NotZero(t, transfer.Transfer.ID)
		require.NotZero(t, transfer.FromEntry.ID)
		require.NotZero(t, transfer.ToEntry.ID)
	}

	updatedAccount1, _ := testQueries.GetAccount(context.Background(), account1.ID)
	updatedAccount2, _ := testQueries.GetAccount(context.Background(), account2.ID)
	require.Equal(t, updatedAccount1.Balance, account1.Balance-(int64(n)*amount))
	require.Equal(t, updatedAccount2.Balance, account2.Balance+(int64(n)*amount))
}
