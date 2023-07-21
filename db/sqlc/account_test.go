package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	"simple-bank/utils"
	"testing"
	"time"
)

func createAccountRandom(t *testing.T) Account {
	a := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	result, err := testQueries.CreateAccount(context.Background(), a)
	idAccount, _ := result.LastInsertId()
	account, _ := testQueries.GetAccount(context.Background(), idAccount)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, a.Owner, account.Owner)
	require.Equal(t, a.Balance, account.Balance)
	require.Equal(t, a.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createAccountRandom(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createAccountRandom(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Owner)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	acc := createAccountRandom(t)

	arg := UpdateAccountParams{
		ID:      acc.ID,
		Balance: acc.Balance,
	}
	result, err := testQueries.UpdateAccount(context.Background(), arg)
	idAcc, _ := result.LastInsertId()
	account2, _ := testQueries.GetAccount(context.Background(), idAcc)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, acc.ID, account2.ID)
	require.Equal(t, acc.Owner, account2.Owner)
	require.Equal(t, acc.Balance, account2.Balance)
	require.Equal(t, acc.Currency, account2.Owner)

	require.WithinDuration(t, acc.UpdatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := createAccountRandom(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.Equal(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {

	for i := 0; i <= 10; i++ {
		createAccountRandom(t)
	}
	limit := 10
	accounts, err := testQueries.ListAccounts(context.Background(), int32(limit))
	require.NoError(t, err)

	for _, account := range accounts {
		require.NotEmpty(t, account)

	}

}
