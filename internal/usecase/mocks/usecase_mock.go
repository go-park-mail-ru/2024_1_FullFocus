// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecase is a generated GoMock package.
package mock_usecase

import (
	context "context"
	reflect "reflect"

	dto "github.com/go-park-mail-ru/2024_1_FullFocus/internal/delivery/dto"
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
func (m *MockAuth) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDBySessionID", ctx, sID)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDBySessionID indicates an expected call of GetUserIDBySessionID.
func (mr *MockAuthMockRecorder) GetUserIDBySessionID(ctx, sID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDBySessionID", reflect.TypeOf((*MockAuth)(nil).GetUserIDBySessionID), ctx, sID)
}

// IsLoggedIn mocks base method.
func (m *MockAuth) IsLoggedIn(ctx context.Context, isID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLoggedIn", ctx, isID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLoggedIn indicates an expected call of IsLoggedIn.
func (mr *MockAuthMockRecorder) IsLoggedIn(ctx, isID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLoggedIn", reflect.TypeOf((*MockAuth)(nil).IsLoggedIn), ctx, isID)
}

// Login mocks base method.
func (m *MockAuth) Login(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), ctx, login, password)
}

// Logout mocks base method.
func (m *MockAuth) Logout(ctx context.Context, sID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", ctx, sID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockAuthMockRecorder) Logout(ctx, sID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockAuth)(nil).Logout), ctx, sID)
}

// Signup mocks base method.
func (m *MockAuth) Signup(ctx context.Context, login, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Signup", ctx, login, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Signup indicates an expected call of Signup.
func (mr *MockAuthMockRecorder) Signup(ctx, login, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Signup", reflect.TypeOf((*MockAuth)(nil).Signup), ctx, login, password)
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
func (m *MockProducts) GetProducts(ctx context.Context, lastID, limit int) ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProducts", ctx, lastID, limit)
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductsMockRecorder) GetProducts(ctx, lastID, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProducts)(nil).GetProducts), ctx, lastID, limit)
}

// MockAvatars is a mock of Avatars interface.
type MockAvatars struct {
	ctrl     *gomock.Controller
	recorder *MockAvatarsMockRecorder
}

// MockAvatarsMockRecorder is the mock recorder for MockAvatars.
type MockAvatarsMockRecorder struct {
	mock *MockAvatars
}

// NewMockAvatars creates a new mock instance.
func NewMockAvatars(ctrl *gomock.Controller) *MockAvatars {
	mock := &MockAvatars{ctrl: ctrl}
	mock.recorder = &MockAvatarsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvatars) EXPECT() *MockAvatarsMockRecorder {
	return m.recorder
}

// DeleteAvatar mocks base method.
func (m *MockAvatars) DeleteAvatar(ctx context.Context, profileID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAvatar", ctx, profileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAvatar indicates an expected call of DeleteAvatar.
func (mr *MockAvatarsMockRecorder) DeleteAvatar(ctx, profileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAvatar", reflect.TypeOf((*MockAvatars)(nil).DeleteAvatar), ctx, profileID)
}

// UploadAvatar mocks base method.
func (m *MockAvatars) UploadAvatar(ctx context.Context, img dto.Image, profileID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", ctx, img, profileID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockAvatarsMockRecorder) UploadAvatar(ctx, img, profileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockAvatars)(nil).UploadAvatar), ctx, img, profileID)
}

// MockOrders is a mock of Orders interface.
type MockOrders struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersMockRecorder
}

// MockOrdersMockRecorder is the mock recorder for MockOrders.
type MockOrdersMockRecorder struct {
	mock *MockOrders
}

// NewMockOrders creates a new mock instance.
func NewMockOrders(ctrl *gomock.Controller) *MockOrders {
	mock := &MockOrders{ctrl: ctrl}
	mock.recorder = &MockOrdersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrders) EXPECT() *MockOrdersMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockOrders) Create(ctx context.Context, input models.CreateOrderInput) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, input)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrdersMockRecorder) Create(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrders)(nil).Create), ctx, input)
}

// Delete mocks base method.
func (m *MockOrders) Delete(ctx context.Context, profileID, orderingID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, profileID, orderingID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrdersMockRecorder) Delete(ctx, profileID, orderingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrders)(nil).Delete), ctx, profileID, orderingID)
}

// GetAllOrders mocks base method.
func (m *MockOrders) GetAllOrders(ctx context.Context, profileID uint) ([]models.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrders", ctx, profileID)
	ret0, _ := ret[0].([]models.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrders indicates an expected call of GetAllOrders.
func (mr *MockOrdersMockRecorder) GetAllOrders(ctx, profileID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrders", reflect.TypeOf((*MockOrders)(nil).GetAllOrders), ctx, profileID)
}

// GetOrderByID mocks base method.
func (m *MockOrders) GetOrderByID(ctx context.Context, profileID, orderingID uint) (models.GetOrderPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", ctx, profileID, orderingID)
	ret0, _ := ret[0].(models.GetOrderPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockOrdersMockRecorder) GetOrderByID(ctx, profileID, orderingID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockOrders)(nil).GetOrderByID), ctx, profileID, orderingID)
}
