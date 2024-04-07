// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	models "github.com/go-park-mail-ru/2024_1_FullFocus/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsers) CreateUser(ctx context.Context, user models.User) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, user)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsersMockRecorder) CreateUser(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsers)(nil).CreateUser), ctx, user)
}

// GetUser mocks base method.
func (m *MockUsers) GetUser(ctx context.Context, login string) (models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx, login)
	ret0, _ := ret[0].(models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUsersMockRecorder) GetUser(ctx, login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUsers)(nil).GetUser), ctx, login)
}

// MockSessions is a mock of Sessions interface.
type MockSessions struct {
	ctrl     *gomock.Controller
	recorder *MockSessionsMockRecorder
}

// MockSessionsMockRecorder is the mock recorder for MockSessions.
type MockSessionsMockRecorder struct {
	mock *MockSessions
}

// NewMockSessions creates a new mock instance.
func NewMockSessions(ctrl *gomock.Controller) *MockSessions {
	mock := &MockSessions{ctrl: ctrl}
	mock.recorder = &MockSessionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSessions) EXPECT() *MockSessionsMockRecorder {
	return m.recorder
}

// CreateSession mocks base method.
func (m *MockSessions) CreateSession(ctx context.Context, userID uint) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSession", ctx, userID)
	ret0, _ := ret[0].(string)
	return ret0
}

// CreateSession indicates an expected call of CreateSession.
func (mr *MockSessionsMockRecorder) CreateSession(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSession", reflect.TypeOf((*MockSessions)(nil).CreateSession), ctx, userID)
}

// DeleteSession mocks base method.
func (m *MockSessions) DeleteSession(ctx context.Context, sID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSession", ctx, sID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockSessionsMockRecorder) DeleteSession(ctx, sID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockSessions)(nil).DeleteSession), ctx, sID)
}

// GetUserIDBySessionID mocks base method.
func (m *MockSessions) GetUserIDBySessionID(ctx context.Context, sID string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserIDBySessionID", ctx, sID)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserIDBySessionID indicates an expected call of GetUserIDBySessionID.
func (mr *MockSessionsMockRecorder) GetUserIDBySessionID(ctx, sID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserIDBySessionID", reflect.TypeOf((*MockSessions)(nil).GetUserIDBySessionID), ctx, sID)
}

// SessionExists mocks base method.
func (m *MockSessions) SessionExists(ctx context.Context, sID string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SessionExists", ctx, sID)
	ret0, _ := ret[0].(bool)
	return ret0
}

// SessionExists indicates an expected call of SessionExists.
func (mr *MockSessionsMockRecorder) SessionExists(ctx, sID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SessionExists", reflect.TypeOf((*MockSessions)(nil).SessionExists), ctx, sID)
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
func (m *MockAvatars) DeleteAvatar(ctx context.Context, imageName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAvatar", ctx, imageName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAvatar indicates an expected call of DeleteAvatar.
func (mr *MockAvatarsMockRecorder) DeleteAvatar(ctx, imageName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAvatar", reflect.TypeOf((*MockAvatars)(nil).DeleteAvatar), ctx, imageName)
}

// UploadAvatar mocks base method.
func (m *MockAvatars) UploadAvatar(ctx context.Context, img models.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadAvatar", ctx, img)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadAvatar indicates an expected call of UploadAvatar.
func (mr *MockAvatarsMockRecorder) UploadAvatar(ctx, img interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadAvatar", reflect.TypeOf((*MockAvatars)(nil).UploadAvatar), ctx, img)
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
func (m *MockOrders) Create(ctx context.Context, userID uint, orderItems []models.OrderItem) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, userID, orderItems)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockOrdersMockRecorder) Create(ctx, userID, orderItems interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockOrders)(nil).Create), ctx, userID, orderItems)
}

// Delete mocks base method.
func (m *MockOrders) Delete(ctx context.Context, orderID uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, orderID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockOrdersMockRecorder) Delete(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockOrders)(nil).Delete), ctx, orderID)
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
func (m *MockOrders) GetOrderByID(ctx context.Context, orderID uint) (models.GetOrderPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByID", ctx, orderID)
	ret0, _ := ret[0].(models.GetOrderPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByID indicates an expected call of GetOrderByID.
func (mr *MockOrdersMockRecorder) GetOrderByID(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByID", reflect.TypeOf((*MockOrders)(nil).GetOrderByID), ctx, orderID)
}

// GetProfileIDByOrderID mocks base method.
func (m *MockOrders) GetProfileIDByOrderID(ctx context.Context, orderID uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileIDByOrderID", ctx, orderID)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileIDByOrderID indicates an expected call of GetProfileIDByOrderID.
func (mr *MockOrdersMockRecorder) GetProfileIDByOrderID(ctx, orderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileIDByOrderID", reflect.TypeOf((*MockOrders)(nil).GetProfileIDByOrderID), ctx, orderID)
}
