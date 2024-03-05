// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// GetUserIDBySessionID mocks base method.
func (m *MockAuth) GetUserIDBySessionID(sID string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDBySessionID", sID)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDBySessionID indicates an expected call of GetUserIDBySessionID.
func (mr *MockAuthMockRecorder) GetUserIDBySessionID(sID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDBySessionID", reflect.TypeOf((*MockAuth)(nil).GetUserIDBySessionID), sID)
}

// IsLoggedIn mocks base method.
func (m *MockAuth) IsLoggedIn(isID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLoggedIn", isID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLoggedIn indicates an expected call of IsLoggedIn.
func (mr *MockAuthMockRecorder) IsLoggedIn(isID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLoggedIn", reflect.TypeOf((*MockAuth)(nil).IsLoggedIn), isID)
}

// Login mocks base method.
func (m *MockAuth) Login(login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), login, password)
}

// Logout mocks base method.
func (m *MockAuth) Logout(sID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", sID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthMockRecorder) Logout(sID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuth)(nil).Logout), sID)
}

// Signup mocks base method.
func (m *MockAuth) Signup(login, password string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthMockRecorder) Signup(login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuth)(nil).Signup), login, password)
}

// MockProducts is a mock of Products interface.
type MockProducts struct {
	ctrl     *gomock.Controller
	recorder *MockProductsMockRecorder
}

// MockProductsMockRecorder is the mock recorder for MockProducts.
type MockProductsMockRecorder struct {
	mock *MockProducts
}

// NewMockProducts creates a new mock instance.
func NewMockProducts(ctrl *gomock.Controller) *MockProducts {
	mock := &MockProducts{ctrl: ctrl}
	mock.recorder = &MockProductsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducts) EXPECT() *MockProductsMockRecorder {
	return m.recorder
}

// GetProducts mocks base method.
func (m *MockProducts) GetProducts(lastID, limit int) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", lastID, limit)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductsMockRecorder) GetProducts(lastID, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProducts)(nil).GetProducts), lastID, limit)
}