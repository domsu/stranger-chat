// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/domsu/stranger-chat/repository (interfaces: IUserRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	model "github.com/domsu/stranger-chat/model"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
)

// MockIUserRepository is a mock of IUserRepository interface
type MockIUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIUserRepositoryMockRecorder
}

// MockIUserRepositoryMockRecorder is the mock recorder for MockIUserRepository
type MockIUserRepositoryMockRecorder struct {
	mock *MockIUserRepository
}

// NewMockIUserRepository creates a new mock instance
func NewMockIUserRepository(ctrl *gomock.Controller) *MockIUserRepository {
	mock := &MockIUserRepository{ctrl: ctrl}
	mock.recorder = &MockIUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIUserRepository) EXPECT() *MockIUserRepositoryMockRecorder {
	return m.recorder
}

// CreateNewUser mocks base method
func (m *MockIUserRepository) CreateNewUser() *model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser")
	ret0, _ := ret[0].(*model.User)
	return ret0
}

// CreateNewUser indicates an expected call of CreateNewUser
func (mr *MockIUserRepositoryMockRecorder) CreateNewUser() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockIUserRepository)(nil).CreateNewUser))
}

// GetActiveUsers mocks base method
func (m *MockIUserRepository) GetActiveUsers() []*model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveUsers")
	ret0, _ := ret[0].([]*model.User)
	return ret0
}

// GetActiveUsers indicates an expected call of GetActiveUsers
func (mr *MockIUserRepositoryMockRecorder) GetActiveUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveUsers", reflect.TypeOf((*MockIUserRepository)(nil).GetActiveUsers))
}

// GetUser mocks base method
func (m *MockIUserRepository) GetUser(arg0 uuid.UUID) *model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(*model.User)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockIUserRepositoryMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockIUserRepository)(nil).GetUser), arg0)
}

// GetUsers mocks base method
func (m *MockIUserRepository) GetUsers() []*model.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*model.User)
	return ret0
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockIUserRepositoryMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockIUserRepository)(nil).GetUsers))
}