package services

import (
	"context"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
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
	Transfer 	db.Transfer 	`json:"transfer"`
	FromAccount db.Account 		`json:"from_account"`
	ToAccount   db.Account 		`json:"to_account"`
	FromEntry   db.Entry 		`json:"from_entry"`
	ToEntry     db.Entry 		`json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// create transfer record, add account entries, update both accounts balance within single transaction
func (s *services) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	// TODO: rewrite execTx to be a generic function to return any type of result so we can get rid of the closure
	err := s.execTx(ctx, func(q *db.Queries) error {
		var err error

		// TODO: result.Transfer is a closure variable
		result.Transfer, err = q.CreateTransfer(ctx, db.CreateTransferParams(arg))
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, db.CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, db.CreateEntryParams{
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

func addMoney(ctx context.Context, q *db.Queries, accountID1 uuid.UUID, amount1 int64, accountID2 uuid.UUID, amount2 int64) (account1, account2 db.Account, err error) {
	account1, err = q.AddAccountBalance(ctx, db.AddAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.AddAccountBalance(ctx, db.AddAccountBalanceParams{
		ID:     	accountID2,
		Amount: 	amount2,
	})
	return
}
