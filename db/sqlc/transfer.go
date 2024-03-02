package db

import (
	"context"

	"github.com/google/uuid"
)

// contains the input parameters for the transfer transaction
type TransferTxParams struct {
	FromAccountID uuid.UUID `json:"from_account_id"`
	ToAccountID   uuid.UUID `json:"to_account_id"`
	Amount        int64 	`json:"amount"`
}

// result of transfer transaction (DTO)
type TransferTxResult struct {
	Transfer 	Transfer 	`json:"transfer"`
	FromAccount Account 	`json:"from_account"`
	ToAccount   Account 	`json:"to_account"`
	FromEntry   Entry 		`json:"from_entry"`
	ToEntry     Entry 		`json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// create transfer record, add account entries, update both accounts balance within single transaction
func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// TODO: rewrite execTx to be a generic function to return any type of result so we can get rid of the closure
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		// TODO: result.Transfer is a closure variable
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams(arg))
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		// update accounts' balance
		if arg.FromAccountID.String() < arg.ToAccountID.String() {
			result.FromAccount, result.ToAccount, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccount, result.FromAccount, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}

		return err
	})

	return result, err
}

func addMoney(ctx context.Context, q *Queries, accountID1 uuid.UUID, amount1 int64, accountID2 uuid.UUID, amount2 int64) (account1 Account, account2 Account, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     	accountID2,
		Amount: 	amount2,
	})
	return
}