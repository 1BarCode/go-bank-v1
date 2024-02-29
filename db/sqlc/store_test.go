package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errCh := make(chan error)
	resCh := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func(errCh chan<- error, resCh chan<- TransferTxResult, amount int64) {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Amount:        amount,
			})
			errCh <- err
			resCh <- result
		}(errCh, resCh, amount)
	}

	// existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errCh
		require.NoError(t, err)

		result := <-resCh
		require.NotEmpty(t, result)

		// check transfer
		xfer := result.Transfer
		require.NotEmpty(t, xfer)
		require.Equal(t, acc1.ID, xfer.FromAccountID)
		require.Equal(t, acc2.ID, xfer.ToAccountID)
		require.Equal(t, amount, xfer.Amount)
		require.NotZero(t, xfer.ID)
		require.NotZero(t, xfer.CreatedAt)
		require.NotZero(t, xfer.UpdatedAt)

		// get transfer from db
		_, err = store.GetTransfer(context.Background(), xfer.ID)
		require.NoError(t, err)


		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, acc1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)
		require.NotZero(t, fromEntry.UpdatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, acc2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)
		require.NotZero(t, toEntry.UpdatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// // TODO: check accounts
		// fromAccount := result.FromAccount
		// require.NotEmpty(t, fromAccount)
		// require.Equal(t, acc1.ID, fromAccount.ID)

		// toAccount := result.ToAccount
		// require.NotEmpty(t, toAccount)
		// require.Equal(t, acc2.ID, toAccount.ID)

		// // check balance
		// diff1 := acc1.Balance - fromAccount.Balance
		// diff2 := toAccount.Balance - acc2.Balance
		// require.Equal(t, diff1, diff2)
		// require.True(t, diff1 > 0)
		// require.True(t, diff1 % amount == 0)

		// k := int(diff1 / amount)
		// require.True(t, k >= 1 && k <= n)
		// require.NotContains(t, existed, k)
		// existed[k] = true
	}

	// // check account balances after all transactions
	// updatedAcc1, err := store.GetAccount(context.Background(), acc1.ID)
	// require.NoError(t, err)

	// updatedAcc2, err := store.GetAccount(context.Background(), acc2.ID)
	// require.NoError(t, err)

	// require.Equal(t, acc1.Balance-int64(n)*amount, updatedAcc1.Balance)
	// require.Equal(t, acc2.Balance+int64(n)*amount, updatedAcc2.Balance)
}