// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/1BarCode/go-bank-v1/services (interfaces: Services)

// Package mockservices is a generated GoMock package.
package mockservices

import (
	context "context"
	reflect "reflect"

	db "github.com/1BarCode/go-bank-v1/db/sqlc"
	services "github.com/1BarCode/go-bank-v1/services"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockServices is a mock of Services interface.
type MockServices struct {
	ctrl     *gomock.Controller
	recorder *MockServicesMockRecorder
}

// MockServicesMockRecorder is the mock recorder for MockServices.
type MockServicesMockRecorder struct {
	mock *MockServices
}

// NewMockServices creates a new mock instance.
func NewMockServices(ctrl *gomock.Controller) *MockServices {
	mock := &MockServices{ctrl: ctrl}
	mock.recorder = &MockServicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServices) EXPECT() *MockServicesMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockServices) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockServicesMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockServices)(nil).CreateAccount), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockServices) DeleteAccount(arg0 context.Context, arg1 uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockServicesMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockServices)(nil).DeleteAccount), arg0, arg1)
}

// GetAccount mocks base method.
func (m *MockServices) GetAccount(arg0 context.Context, arg1 uuid.UUID) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockServicesMockRecorder) GetAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockServices)(nil).GetAccount), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockServices) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockServicesMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockServices)(nil).ListAccounts), arg0, arg1)
}

// TransferTx mocks base method.
func (m *MockServices) TransferTx(arg0 context.Context, arg1 services.TransferTxParams) (services.TransferTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferTx", arg0, arg1)
	ret0, _ := ret[0].(services.TransferTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransferTx indicates an expected call of TransferTx.
func (mr *MockServicesMockRecorder) TransferTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferTx", reflect.TypeOf((*MockServices)(nil).TransferTx), arg0, arg1)
}
