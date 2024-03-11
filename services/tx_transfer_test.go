package services

import (
	"context"
	"testing"
	"time"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/1BarCode/go-bank-v1/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, fromAcc, toAcc db.Account) db.Transfer {
	arg := db.CreateTransferParams{
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

	arg := db.ListTransfersParams{
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

func TestTransferTx(t *testing.T) {
	store := db.NewStore(testDB) // need store to call TransferTx
	services := NewServices(store)
	
	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errCh := make(chan error)
	resCh := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func(errCh chan<- error, resCh chan<- TransferTxResult, amount int64) {
			result, err := services.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: acc1.ID,
				ToAccountID:   acc2.ID,
				Amount:        amount,
			})
			errCh <- err
			resCh <- result
		}(errCh, resCh, amount)
	}

	existed := make(map[int]bool)

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

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, acc1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, acc2.ID, toAccount.ID)

		// check balance
		diff1 := acc1.Balance - fromAccount.Balance // amount deducted from acc1
		diff2 := toAccount.Balance - acc2.Balance // amount added to acc2
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1 % amount == 0)

		k := int(diff1 / amount) // number of transactions
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check account balances after all transactions
	updatedAcc1, err := store.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	updatedAcc2, err := store.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)

	require.Equal(t, acc1.Balance-int64(n)*amount, updatedAcc1.Balance)
	require.Equal(t, acc2.Balance+int64(n)*amount, updatedAcc2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := db.NewStore(testDB) // need store to call TransferTx
	services := NewServices(store)

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 10
	amount := int64(10)
	errCh := make(chan error)

	// half of the transactions will transfer from acc1 to acc2
	// the other half will transfer from acc2 to acc1
	for i := 0; i < n; i++ {
		fromAccountID := acc1.ID
		toAccountID := acc2.ID
		if i % 2 == 1 {
			fromAccountID = acc2.ID
			toAccountID = acc1.ID
		}

		go func(errCh chan<- error, amount int64, fromAccountId, toAccountId uuid.UUID) {
			_, err := services.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: 	fromAccountId,
				ToAccountID:  	toAccountId,
				Amount:       	amount,
			})
			errCh <- err
		}(errCh, amount, fromAccountID, toAccountID)
	}

	for i := 0; i < n; i++ {
		err := <-errCh
		require.NoError(t, err)
	}

	// check account balances after all transactions
	updatedAcc1, err := store.GetAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	updatedAcc2, err := store.GetAccount(context.Background(), acc2.ID)
	require.NoError(t, err)

	require.Equal(t, acc1.Balance, updatedAcc1.Balance)
	require.Equal(t, acc2.Balance, updatedAcc2.Balance)
}

// Note function does not start with "Test" so it will not be run by "go test" as a unit test
func createRandomAccount(t *testing.T) db.Account {
	user := createRandomUser(t)

	arg := db.CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)

	return account
}

func createRandomUser(t *testing.T) db.User {
	arg := db.CreateUserParams{
		Username: util.RandomOwner(),
		Email:   util.RandomEmail(),
		HashedPassword: "password",
		FirstName: util.RandomOwner(),
		LastName: util.RandomOwner(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}