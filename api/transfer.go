package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/1BarCode/go-bank-v1/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type transferRequest struct {
	FromAccountID string `json:"from_account_id" binding:"required,uuid"`
	ToAccountID   string `json:"to_account_id" binding:"required,uuid"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

// TODO: write unit test for this
func (s *Server) createTransfer(ctx *gin.Context) {
	var req transferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// convert account IDs from string to uuid
	fromAccountID, _ := uuid.Parse(req.FromAccountID)
	toAccountID, _ := uuid.Parse(req.ToAccountID)

	// validate accounts
	if _, valid := s.validateAccountForTransfer(ctx, fromAccountID, req.Currency); !valid {
		return
	}

	if _, valid := s.validateAccountForTransfer(ctx, toAccountID, req.Currency); !valid {
		return
	}

	arg := services.TransferTxParams{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        req.Amount,
	}

	result, err := s.services.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, intServerErrorResponse())
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (s *Server) validateAccountForTransfer(ctx *gin.Context, accountID uuid.UUID, currency string) (db.Account, bool) {
	account, err := s.services.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return account, false
		}

		ctx.JSON(http.StatusInternalServerError, intServerErrorResponse())
		return account, false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%v] currency mismatch: %s vs given %s", account.ID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return account, false
	}

	return account, true
}