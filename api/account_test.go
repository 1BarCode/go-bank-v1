package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	mockservices "github.com/1BarCode/go-bank-v1/services/mock"
	"github.com/1BarCode/go-bank-v1/util"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {
	testAcc := randomAccount()

	testCases := []struct {
		name       string
		accountID  uuid.UUID
		buildStubs func(services *mockservices.MockServices)
		checkRes   func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			accountID: testAcc.ID,
			buildStubs: func(services *mockservices.MockServices) {
				services.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(testAcc.ID)).
					Times(1).
					Return(testAcc, nil)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.Equal(t, "application/json; charset=utf-8", recorder.Header().Get("Content-Type"))
				requireBodyMatchAccount(t, recorder.Body, testAcc)
			},
		},
		{
			name:      "Not Found",
			accountID: testAcc.ID,
			buildStubs: func(services *mockservices.MockServices) {
				services.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(testAcc.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
			
		},
		{
			name:      "Internal Error",
			accountID: testAcc.ID,
			buildStubs: func(services *mockservices.MockServices) {
				services.EXPECT().
					GetAccount(gomock.Any(), gomock.Eq(testAcc.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrConnDone)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{	
			name:      "Invalid ID",
			accountID: 	uuid.Nil,
			buildStubs: func(services *mockservices.MockServices) {
				services.EXPECT().
					GetAccount(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkRes: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockServices := mockservices.NewMockServices(ctrl)
			// build stub
			tc.buildStubs(mockServices)

			// create test server and send request
			server := NewServer(mockServices)
			recorder := httptest.NewRecorder()

			url := "/v1/accounts/" + tc.accountID.String()
			if (tc.name == "Invalid ID") {
				url = "/v1/accounts/" + "invalidID"
			}
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			// check response
			tc.checkRes(t, recorder)
		})
	}
}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomUuid(),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}