// Code generated by MockGen. DO NOT EDIT.
// Source: internal/users/usecase/usecase.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "github.com/DoWithLogic/golang-clean-architecture/internal/users/entities"
	gomock "go.uber.org/mock/gomock"
)

// MockUsecase is a mock of Usecase interface.
type MockUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockUsecaseMockRecorder
}

// MockUsecaseMockRecorder is the mock recorder for MockUsecase.
type MockUsecaseMockRecorder struct {
	mock *MockUsecase
}

// NewMockUsecase creates a new mock instance.
func NewMockUsecase(ctrl *gomock.Controller) *MockUsecase {
	mock := &MockUsecase{ctrl: ctrl}
	mock.recorder = &MockUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsecase) EXPECT() *MockUsecaseMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsecase) CreateUser(ctx context.Context, user entities.CreateUser) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsecaseMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsecase)(nil).CreateUser), ctx, user)
}

// UpdateUser mocks base method.
func (m *MockUsecase) UpdateUser(ctx context.Context, updateData entities.UpdateUsers) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, updateData)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUsecaseMockRecorder) UpdateUser(ctx, updateData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUsecase)(nil).UpdateUser), ctx, updateData)
}

// UpdateUserStatus mocks base method.
func (m *MockUsecase) UpdateUserStatus(ctx context.Context, req entities.UpdateUserStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserStatus", ctx, req)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserStatus indicates an expected call of UpdateUserStatus.
func (mr *MockUsecaseMockRecorder) UpdateUserStatus(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserStatus", reflect.TypeOf((*MockUsecase)(nil).UpdateUserStatus), ctx, req)
}
