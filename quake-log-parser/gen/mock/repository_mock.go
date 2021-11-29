// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Ralphbaer/cloudwalk/quake-log-parser/repository (interfaces: QuakeLogRepository)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	os "os"
	reflect "reflect"
)

// MockQuakeLogRepository is a mock of QuakeLogRepository interface
type MockQuakeLogRepository struct {
	ctrl     *gomock.Controller
	recorder *MockQuakeLogRepositoryMockRecorder
}

// MockQuakeLogRepositoryMockRecorder is the mock recorder for MockQuakeLogRepository
type MockQuakeLogRepositoryMockRecorder struct {
	mock *MockQuakeLogRepository
}

// NewMockQuakeLogRepository creates a new mock instance
func NewMockQuakeLogRepository(ctrl *gomock.Controller) *MockQuakeLogRepository {
	mock := &MockQuakeLogRepository{ctrl: ctrl}
	mock.recorder = &MockQuakeLogRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuakeLogRepository) EXPECT() *MockQuakeLogRepositoryMockRecorder {
	return m.recorder
}

// GetFile mocks base method
func (m *MockQuakeLogRepository) GetFile(arg0 context.Context) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", arg0)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockQuakeLogRepositoryMockRecorder) GetFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockQuakeLogRepository)(nil).GetFile), arg0)
}
