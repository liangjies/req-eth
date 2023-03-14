// Code generated by MockGen. DO NOT EDIT.
// Source: websocket.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	eth "github.com/liangjies/req-eth/eth"
	jsonrpc "github.com/liangjies/req-eth/jsonrpc"
	websocket "github.com/liangjies/req-eth/node/websocket"
	gomock "github.com/golang/mock/gomock"
)

// MockConnection is a mock of Connection interface.
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionMockRecorder
}

// MockConnectionMockRecorder is the mock recorder for MockConnection.
type MockConnectionMockRecorder struct {
	mock *MockConnection
}

// NewMockConnection creates a new mock instance.
func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &MockConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnection) EXPECT() *MockConnectionMockRecorder {
	return m.recorder
}

// BlockByHash mocks base method.
func (m *MockConnection) BlockByHash(ctx context.Context, hash string, full bool) (*eth.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHash", ctx, hash, full)
	ret0, _ := ret[0].(*eth.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHash indicates an expected call of BlockByHash.
func (mr *MockConnectionMockRecorder) BlockByHash(ctx, hash, full interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHash", reflect.TypeOf((*MockConnection)(nil).BlockByHash), ctx, hash, full)
}

// BlockByNumber mocks base method.
func (m *MockConnection) BlockByNumber(ctx context.Context, number uint64, full bool) (*eth.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByNumber", ctx, number, full)
	ret0, _ := ret[0].(*eth.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByNumber indicates an expected call of BlockByNumber.
func (mr *MockConnectionMockRecorder) BlockByNumber(ctx, number, full interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByNumber", reflect.TypeOf((*MockConnection)(nil).BlockByNumber), ctx, number, full)
}

// BlockNumber mocks base method.
func (m *MockConnection) BlockNumber(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber indicates an expected call of BlockNumber.
func (mr *MockConnectionMockRecorder) BlockNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockConnection)(nil).BlockNumber), ctx)
}

// GetLogs mocks base method.
func (m *MockConnection) GetLogs(ctx context.Context, filter eth.LogFilter) ([]eth.Log, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogs", ctx, filter)
	ret0, _ := ret[0].([]eth.Log)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLogs indicates an expected call of GetLogs.
func (mr *MockConnectionMockRecorder) GetLogs(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogs", reflect.TypeOf((*MockConnection)(nil).GetLogs), ctx, filter)
}

// NewHeads mocks base method.
func (m *MockConnection) NewHeads(ctx context.Context) (websocket.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHeads", ctx)
	ret0, _ := ret[0].(websocket.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewHeads indicates an expected call of NewHeads.
func (mr *MockConnectionMockRecorder) NewHeads(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHeads", reflect.TypeOf((*MockConnection)(nil).NewHeads), ctx)
}

// NewPendingTransaction mocks base method.
func (m *MockConnection) NewPendingTransaction(ctx context.Context, full ...bool) (websocket.Subscription, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range full {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewPendingTransaction", varargs...)
	ret0, _ := ret[0].(websocket.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPendingTransaction indicates an expected call of NewPendingTransaction.
func (mr *MockConnectionMockRecorder) NewPendingTransaction(ctx interface{}, full ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, full...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPendingTransaction", reflect.TypeOf((*MockConnection)(nil).NewPendingTransaction), varargs...)
}

// Request mocks base method.
func (m *MockConnection) Request(ctx context.Context, r *jsonrpc.Request) (*jsonrpc.RawResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", ctx, r)
	ret0, _ := ret[0].(*jsonrpc.RawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockConnectionMockRecorder) Request(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockConnection)(nil).Request), ctx, r)
}

// Subscribe mocks base method.
func (m *MockConnection) Subscribe(ctx context.Context, r *jsonrpc.Request) (websocket.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", ctx, r)
	ret0, _ := ret[0].(websocket.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockConnectionMockRecorder) Subscribe(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockConnection)(nil).Subscribe), ctx, r)
}

// TransactionByHash mocks base method.
func (m *MockConnection) TransactionByHash(ctx context.Context, hash string) (*eth.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, hash)
	ret0, _ := ret[0].(*eth.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByHash indicates an expected call of TransactionByHash.
func (mr *MockConnectionMockRecorder) TransactionByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockConnection)(nil).TransactionByHash), ctx, hash)
}

// TransactionReceipt mocks base method.
func (m *MockConnection) TransactionReceipt(ctx context.Context, hash string) (*eth.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, hash)
	ret0, _ := ret[0].(*eth.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt indicates an expected call of TransactionReceipt.
func (mr *MockConnectionMockRecorder) TransactionReceipt(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockConnection)(nil).TransactionReceipt), ctx, hash)
}

// URL mocks base method.
func (m *MockConnection) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL.
func (mr *MockConnectionMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockConnection)(nil).URL))
}