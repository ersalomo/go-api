package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	a := CreateAccountParams{
		Owner:    "Ersalomo",
		Balance:  1000,
		Currency: "IDK",
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

}
