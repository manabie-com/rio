// Code generated by MockGen. DO NOT EDIT.
// Source: stub_store.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	rio "github.com/hungdv136/rio"
)

// MockStubStore is a mock of StubStore interface.
type MockStubStore struct {
	ctrl     *gomock.Controller
	recorder *MockStubStoreMockRecorder
}

// MockStubStoreMockRecorder is the mock recorder for MockStubStore.
type MockStubStoreMockRecorder struct {
	mock *MockStubStore
}

// NewMockStubStore creates a new mock instance.
func NewMockStubStore(ctrl *gomock.Controller) *MockStubStore {
	mock := &MockStubStore{ctrl: ctrl}
	mock.recorder = &MockStubStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStubStore) EXPECT() *MockStubStoreMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockStubStore) Create(ctx context.Context, stubs ...*rio.Stub) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range stubs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockStubStoreMockRecorder) Create(ctx interface{}, stubs ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, stubs...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockStubStore)(nil).Create), varargs...)
}

// CreateIncomingRequest mocks base method.
func (m *MockStubStore) CreateIncomingRequest(ctx context.Context, r *rio.IncomingRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIncomingRequest", ctx, r)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIncomingRequest indicates an expected call of CreateIncomingRequest.
func (mr *MockStubStoreMockRecorder) CreateIncomingRequest(ctx, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIncomingRequest", reflect.TypeOf((*MockStubStore)(nil).CreateIncomingRequest), ctx, r)
}

// CreateProto mocks base method.
func (m *MockStubStore) CreateProto(ctx context.Context, protos ...*rio.Proto) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx}
	for _, a := range protos {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProto", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProto indicates an expected call of CreateProto.
func (mr *MockStubStoreMockRecorder) CreateProto(ctx interface{}, protos ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx}, protos...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProto", reflect.TypeOf((*MockStubStore)(nil).CreateProto), varargs...)
}

// Delete mocks base method.
func (m *MockStubStore) Delete(ctx context.Context, id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStubStoreMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStubStore)(nil).Delete), ctx, id)
}

// GetAll mocks base method.
func (m *MockStubStore) GetAll(ctx context.Context, namespace string) ([]*rio.Stub, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, namespace)
	ret0, _ := ret[0].([]*rio.Stub)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockStubStoreMockRecorder) GetAll(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockStubStore)(nil).GetAll), ctx, namespace)
}

// GetIncomingRequests mocks base method.
func (m *MockStubStore) GetIncomingRequests(ctx context.Context, option *rio.IncomingQueryOption) ([]*rio.IncomingRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIncomingRequests", ctx, option)
	ret0, _ := ret[0].([]*rio.IncomingRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIncomingRequests indicates an expected call of GetIncomingRequests.
func (mr *MockStubStoreMockRecorder) GetIncomingRequests(ctx, option interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIncomingRequests", reflect.TypeOf((*MockStubStore)(nil).GetIncomingRequests), ctx, option)
}

// GetProtos mocks base method.
func (m *MockStubStore) GetProtos(ctx context.Context) ([]*rio.Proto, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProtos", ctx)
	ret0, _ := ret[0].([]*rio.Proto)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProtos indicates an expected call of GetProtos.
func (mr *MockStubStoreMockRecorder) GetProtos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProtos", reflect.TypeOf((*MockStubStore)(nil).GetProtos), ctx)
}

// Reset mocks base method.
func (m *MockStubStore) Reset(ctx context.Context, option *rio.ResetQueryOption) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset", ctx, option)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockStubStoreMockRecorder) Reset(ctx, option interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockStubStore)(nil).Reset), ctx, option)
}

// MockStatusStore is a mock of StatusStore interface.
type MockStatusStore struct {
	ctrl     *gomock.Controller
	recorder *MockStatusStoreMockRecorder
}

// MockStatusStoreMockRecorder is the mock recorder for MockStatusStore.
type MockStatusStoreMockRecorder struct {
	mock *MockStatusStore
}

// NewMockStatusStore creates a new mock instance.
func NewMockStatusStore(ctrl *gomock.Controller) *MockStatusStore {
	mock := &MockStatusStore{ctrl: ctrl}
	mock.recorder = &MockStatusStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatusStore) EXPECT() *MockStatusStoreMockRecorder {
	return m.recorder
}

// GetLastUpdatedProto mocks base method.
func (m *MockStatusStore) GetLastUpdatedProto(ctx context.Context) (*rio.LastUpdatedRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastUpdatedProto", ctx)
	ret0, _ := ret[0].(*rio.LastUpdatedRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastUpdatedProto indicates an expected call of GetLastUpdatedProto.
func (mr *MockStatusStoreMockRecorder) GetLastUpdatedProto(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastUpdatedProto", reflect.TypeOf((*MockStatusStore)(nil).GetLastUpdatedProto), ctx)
}

// GetLastUpdatedStub mocks base method.
func (m *MockStatusStore) GetLastUpdatedStub(ctx context.Context, namespace string) (*rio.LastUpdatedRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLastUpdatedStub", ctx, namespace)
	ret0, _ := ret[0].(*rio.LastUpdatedRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLastUpdatedStub indicates an expected call of GetLastUpdatedStub.
func (mr *MockStatusStoreMockRecorder) GetLastUpdatedStub(ctx, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLastUpdatedStub", reflect.TypeOf((*MockStatusStore)(nil).GetLastUpdatedStub), ctx, namespace)
}
