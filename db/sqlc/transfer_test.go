package db

import (
	"context"
	"testing"
	"time"

	"github.com/1BarCode/go-bank-v1/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAcc, toAcc Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: 	fromAcc.ID,
		ToAccountID: 	toAcc.ID,
		Amount:			util.RandomMoney(),
	}

	xfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, xfer)

	require.Equal(t, arg.FromAccountID, xfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, xfer.ToAccountID)
	require.Equal(t, arg.Amount, xfer.Amount)

	require.NotZero(t, xfer.ID)
	require.NotZero(t, xfer.CreatedAt)
	require.NotZero(t, xfer.UpdatedAt)

	return xfer
}

func TestCreateTransfer(t *testing.T) {
	fromAcc := createRandomAccount(t)
	toAcc := createRandomAccount(t)
	createRandomTransfer(t, fromAcc, toAcc)
}

func TestGetTransfer(t *testing.T) {
	fromAcc := createRandomAccount(t)
	toAcc := createRandomAccount(t)
	xfer1 := createRandomTransfer(t, fromAcc, toAcc)
	
	xfer2, err := testQueries.GetTransfer(context.Background(), xfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, xfer2)

	// make sure all properties match
	require.Equal(t, xfer1.ID, xfer2.ID)
	require.Equal(t, xfer1.Amount , xfer2.Amount )
	require.Equal(t, xfer1.FromAccountID , xfer2.FromAccountID )
	require.Equal(t, xfer1.ToAccountID , xfer2.ToAccountID )
	require.Equal(t, xfer1.Amount , xfer2.Amount )

	// make sure time within a sec
	require.WithinDuration(t, xfer1.CreatedAt, xfer2.CreatedAt, time.Second)
	require.WithinDuration(t, xfer1.UpdatedAt, xfer2.UpdatedAt, time.Second)
}

func TestListTransfer(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	for i := 0; i < 5; i++ {
		createRandomTransfer(t, acc1, acc2)
		createRandomTransfer(t, acc2, acc1)
	}

	arg := ListTransfersParams{
		FromAccountID: acc1.ID,
		ToAccountID: acc1.ID,
		Limit: 5,
		Offset: 5,
	}
	
	xfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, xfers, 5)

	for _, xfer := range xfers {
		require.NotEmpty(t, xfer)
		require.True(t, xfer.FromAccountID == acc1.ID || xfer.ToAccountID == acc1.ID)
	}
}