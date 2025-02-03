package db

import (
	"context"
	"testing"

	"github.com/Pokala15/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
		Status:   AccountStatusACTIVE,
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	arg := UpdateAccountBalanceParams{
		ID:     account1.ID,
		Amount: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccountBalance(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Balance+arg.Amount, account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	updatedAccount, err := testQueries.UpdateAccountStatus(
		context.Background(),
		UpdateAccountStatusParams{
			ID:     account1.ID,
			Status: AccountStatusCLOSED,
		},
	)
	require.NoError(t, err)
	require.Equal(t, updatedAccount.Status, AccountStatusCLOSED)
}
