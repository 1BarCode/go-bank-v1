package services

import (
	"context"
	"testing"

	mockdb "github.com/1BarCode/go-bank-v1/db/mock"
	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	"github.com/1BarCode/go-bank-v1/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccountService(t *testing.T) {
	testAcc := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := mockdb.NewMockStore(ctrl)
	
	// build stub
	mockStore.EXPECT().
	GetAccount(gomock.Any(), gomock.Eq(testAcc.ID)).
	Times(1).
	Return(testAcc, nil)

	// create test services
	services := NewServices(mockStore)

	// call the method
	account, err := services.GetAccount(context.Background(), testAcc.ID)
	require.NoError(t, err)
	require.Equal(t, testAcc, account)
	require.Equal(t, testAcc.ID, account.ID)
	require.Equal(t, testAcc.Owner, account.Owner)
	require.Equal(t, testAcc.Balance, account.Balance)
	require.Equal(t, testAcc.Currency, account.Currency)
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomUuid(),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}