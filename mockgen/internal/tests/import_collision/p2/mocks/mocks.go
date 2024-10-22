// Code generated by MockGen. DO NOT EDIT.
// Source: go.uber.org/mock/mockgen/internal/tests/import_collision/p2 (interfaces: Mything)
//
// Generated by this command:
//
//	mockgen -destination=mocks/mocks.go -package=internalpackage . Mything
//

// Package internalpackage is a generated GoMock package.
package internalpackage

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	internalpackage "go.uber.org/mock/mockgen/internal/tests/import_collision/internalpackage"
)

// MockMything is a mock of Mything interface.
type MockMything struct {
	ctrl     *gomock.Controller
	recorder *MockMythingMockRecorder
	isgomock struct{}
}

// MockMythingMockRecorder is the mock recorder for MockMything.
type MockMythingMockRecorder struct {
	mock *MockMything
}

// NewMockMything creates a new mock instance.
func NewMockMything(ctrl *gomock.Controller) *MockMything {
	mock := &MockMything{ctrl: ctrl}
	mock.recorder = &MockMythingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMything) EXPECT() *MockMythingMockRecorder {
	return m.recorder
}

// DoThat mocks base method.
func (m *MockMything) DoThat(arg0 int) internalpackage.FooExported {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DoThat", arg0)
	ret0, _ := ret[0].(internalpackage.FooExported)
	return ret0
}

// DoThat indicates an expected call of DoThat.
func (mr *MockMythingMockRecorder) DoThat(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoThat", reflect.TypeOf((*MockMything)(nil).DoThat), arg0)
}
