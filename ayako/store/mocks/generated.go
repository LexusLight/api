// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_store is a generated GoMock package.
package mock_store

import (
	context "context"
	entity "github.com/deissh/rl/ayako/entity"
	store "github.com/deissh/rl/ayako/store"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockStore is a mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// OAuth mocks base method
func (m *MockStore) OAuth() store.OAuth {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OAuth")
	ret0, _ := ret[0].(store.OAuth)
	return ret0
}

// OAuth indicates an expected call of OAuth
func (mr *MockStoreMockRecorder) OAuth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OAuth", reflect.TypeOf((*MockStore)(nil).OAuth))
}

// Beatmap mocks base method
func (m *MockStore) Beatmap() store.Beatmap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Beatmap")
	ret0, _ := ret[0].(store.Beatmap)
	return ret0
}

// Beatmap indicates an expected call of Beatmap
func (mr *MockStoreMockRecorder) Beatmap() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Beatmap", reflect.TypeOf((*MockStore)(nil).Beatmap))
}

// BeatmapSet mocks base method
func (m *MockStore) BeatmapSet() store.BeatmapSet {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeatmapSet")
	ret0, _ := ret[0].(store.BeatmapSet)
	return ret0
}

// BeatmapSet indicates an expected call of BeatmapSet
func (mr *MockStoreMockRecorder) BeatmapSet() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeatmapSet", reflect.TypeOf((*MockStore)(nil).BeatmapSet))
}

// User mocks base method
func (m *MockStore) User() store.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "User")
	ret0, _ := ret[0].(store.User)
	return ret0
}

// User indicates an expected call of User
func (mr *MockStoreMockRecorder) User() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "User", reflect.TypeOf((*MockStore)(nil).User))
}

// Friend mocks base method
func (m *MockStore) Friend() store.Friend {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Friend")
	ret0, _ := ret[0].(store.Friend)
	return ret0
}

// Friend indicates an expected call of Friend
func (mr *MockStoreMockRecorder) Friend() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Friend", reflect.TypeOf((*MockStore)(nil).Friend))
}

// MockOAuth is a mock of OAuth interface
type MockOAuth struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthMockRecorder
}

// MockOAuthMockRecorder is the mock recorder for MockOAuth
type MockOAuthMockRecorder struct {
	mock *MockOAuth
}

// NewMockOAuth creates a new mock instance
func NewMockOAuth(ctrl *gomock.Controller) *MockOAuth {
	mock := &MockOAuth{ctrl: ctrl}
	mock.recorder = &MockOAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOAuth) EXPECT() *MockOAuthMockRecorder {
	return m.recorder
}

// CreateClient mocks base method
func (m *MockOAuth) CreateClient(ctx context.Context, name, redirect string) (*entity.OAuthClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateClient", ctx, name, redirect)
	ret0, _ := ret[0].(*entity.OAuthClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateClient indicates an expected call of CreateClient
func (mr *MockOAuthMockRecorder) CreateClient(ctx, name, redirect interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateClient", reflect.TypeOf((*MockOAuth)(nil).CreateClient), ctx, name, redirect)
}

// GetClient mocks base method
func (m *MockOAuth) GetClient(ctx context.Context, id uint, secret string) (*entity.OAuthClient, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClient", ctx, id, secret)
	ret0, _ := ret[0].(*entity.OAuthClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient
func (mr *MockOAuthMockRecorder) GetClient(ctx, id, secret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockOAuth)(nil).GetClient), ctx, id, secret)
}

// CreateToken mocks base method
func (m *MockOAuth) CreateToken(ctx context.Context, userId, clientID uint, clientSecret, scopes string) (*entity.OAuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", ctx, userId, clientID, clientSecret, scopes)
	ret0, _ := ret[0].(*entity.OAuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken
func (mr *MockOAuthMockRecorder) CreateToken(ctx, userId, clientID, clientSecret, scopes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockOAuth)(nil).CreateToken), ctx, userId, clientID, clientSecret, scopes)
}

// RevokeToken mocks base method
func (m *MockOAuth) RevokeToken(ctx context.Context, userId uint, accessToken string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeToken", ctx, userId, accessToken)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeToken indicates an expected call of RevokeToken
func (mr *MockOAuthMockRecorder) RevokeToken(ctx, userId, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeToken", reflect.TypeOf((*MockOAuth)(nil).RevokeToken), ctx, userId, accessToken)
}

// RefreshToken mocks base method
func (m *MockOAuth) RefreshToken(ctx context.Context, refreshToken string, clientID uint, clientSecret string) (*entity.OAuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshToken", ctx, refreshToken, clientID, clientSecret)
	ret0, _ := ret[0].(*entity.OAuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshToken indicates an expected call of RefreshToken
func (mr *MockOAuthMockRecorder) RefreshToken(ctx, refreshToken, clientID, clientSecret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshToken", reflect.TypeOf((*MockOAuth)(nil).RefreshToken), ctx, refreshToken, clientID, clientSecret)
}

// ValidateToken mocks base method
func (m *MockOAuth) ValidateToken(ctx context.Context, accessToken string) (*entity.OAuthToken, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateToken", ctx, accessToken)
	ret0, _ := ret[0].(*entity.OAuthToken)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateToken indicates an expected call of ValidateToken
func (mr *MockOAuthMockRecorder) ValidateToken(ctx, accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateToken", reflect.TypeOf((*MockOAuth)(nil).ValidateToken), ctx, accessToken)
}

// MockBeatmap is a mock of Beatmap interface
type MockBeatmap struct {
	ctrl     *gomock.Controller
	recorder *MockBeatmapMockRecorder
}

// MockBeatmapMockRecorder is the mock recorder for MockBeatmap
type MockBeatmapMockRecorder struct {
	mock *MockBeatmap
}

// NewMockBeatmap creates a new mock instance
func NewMockBeatmap(ctrl *gomock.Controller) *MockBeatmap {
	mock := &MockBeatmap{ctrl: ctrl}
	mock.recorder = &MockBeatmapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeatmap) EXPECT() *MockBeatmapMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockBeatmap) Get(ctx context.Context, id uint) (*entity.SingleBeatmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*entity.SingleBeatmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBeatmapMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBeatmap)(nil).Get), ctx, id)
}

// GetBySetId mocks base method
func (m *MockBeatmap) GetBySetId(ctx context.Context, beatmapsetId uint) []entity.Beatmap {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBySetId", ctx, beatmapsetId)
	ret0, _ := ret[0].([]entity.Beatmap)
	return ret0
}

// GetBySetId indicates an expected call of GetBySetId
func (mr *MockBeatmapMockRecorder) GetBySetId(ctx, beatmapsetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBySetId", reflect.TypeOf((*MockBeatmap)(nil).GetBySetId), ctx, beatmapsetId)
}

// Create mocks base method
func (m *MockBeatmap) Create(ctx context.Context, from interface{}) (*entity.Beatmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, from)
	ret0, _ := ret[0].(*entity.Beatmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBeatmapMockRecorder) Create(ctx, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBeatmap)(nil).Create), ctx, from)
}

// CreateBatch mocks base method
func (m *MockBeatmap) CreateBatch(ctx context.Context, from interface{}) (*[]entity.Beatmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBatch", ctx, from)
	ret0, _ := ret[0].(*[]entity.Beatmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateBatch indicates an expected call of CreateBatch
func (mr *MockBeatmapMockRecorder) CreateBatch(ctx, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBatch", reflect.TypeOf((*MockBeatmap)(nil).CreateBatch), ctx, from)
}

// Update mocks base method
func (m *MockBeatmap) Update(ctx context.Context, id uint, from interface{}) (*entity.Beatmap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, from)
	ret0, _ := ret[0].(*entity.Beatmap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockBeatmapMockRecorder) Update(ctx, id, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBeatmap)(nil).Update), ctx, id, from)
}

// Delete mocks base method
func (m *MockBeatmap) Delete(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBeatmapMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBeatmap)(nil).Delete), ctx, id)
}

// MockBeatmapSet is a mock of BeatmapSet interface
type MockBeatmapSet struct {
	ctrl     *gomock.Controller
	recorder *MockBeatmapSetMockRecorder
}

// MockBeatmapSetMockRecorder is the mock recorder for MockBeatmapSet
type MockBeatmapSetMockRecorder struct {
	mock *MockBeatmapSet
}

// NewMockBeatmapSet creates a new mock instance
func NewMockBeatmapSet(ctrl *gomock.Controller) *MockBeatmapSet {
	mock := &MockBeatmapSet{ctrl: ctrl}
	mock.recorder = &MockBeatmapSetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBeatmapSet) EXPECT() *MockBeatmapSetMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockBeatmapSet) Get(ctx context.Context, id uint) (*entity.BeatmapSetFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*entity.BeatmapSetFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockBeatmapSetMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBeatmapSet)(nil).Get), ctx, id)
}

// GetAll mocks base method
func (m *MockBeatmapSet) GetAll(ctx context.Context, page, limit int) (*[]entity.BeatmapSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, page, limit)
	ret0, _ := ret[0].(*[]entity.BeatmapSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockBeatmapSetMockRecorder) GetAll(ctx, page, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockBeatmapSet)(nil).GetAll), ctx, page, limit)
}

// ComputeFields mocks base method
func (m *MockBeatmapSet) ComputeFields(ctx context.Context, set entity.BeatmapSetFull) (*entity.BeatmapSetFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeFields", ctx, set)
	ret0, _ := ret[0].(*entity.BeatmapSetFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeFields indicates an expected call of ComputeFields
func (mr *MockBeatmapSetMockRecorder) ComputeFields(ctx, set interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeFields", reflect.TypeOf((*MockBeatmapSet)(nil).ComputeFields), ctx, set)
}

// SetFavourite mocks base method
func (m *MockBeatmapSet) SetFavourite(ctx context.Context, userId, id uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFavourite", ctx, userId, id)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetFavourite indicates an expected call of SetFavourite
func (mr *MockBeatmapSetMockRecorder) SetFavourite(ctx, userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFavourite", reflect.TypeOf((*MockBeatmapSet)(nil).SetFavourite), ctx, userId, id)
}

// SetUnFavourite mocks base method
func (m *MockBeatmapSet) SetUnFavourite(ctx context.Context, userId, id uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetUnFavourite", ctx, userId, id)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetUnFavourite indicates an expected call of SetUnFavourite
func (mr *MockBeatmapSetMockRecorder) SetUnFavourite(ctx, userId, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUnFavourite", reflect.TypeOf((*MockBeatmapSet)(nil).SetUnFavourite), ctx, userId, id)
}

// GetLatestId mocks base method
func (m *MockBeatmapSet) GetLatestId(ctx context.Context) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestId", ctx)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestId indicates an expected call of GetLatestId
func (mr *MockBeatmapSetMockRecorder) GetLatestId(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestId", reflect.TypeOf((*MockBeatmapSet)(nil).GetLatestId), ctx)
}

// GetIdsForUpdate mocks base method
func (m *MockBeatmapSet) GetIdsForUpdate(ctx context.Context, limit int) ([]uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIdsForUpdate", ctx, limit)
	ret0, _ := ret[0].([]uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIdsForUpdate indicates an expected call of GetIdsForUpdate
func (mr *MockBeatmapSetMockRecorder) GetIdsForUpdate(ctx, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIdsForUpdate", reflect.TypeOf((*MockBeatmapSet)(nil).GetIdsForUpdate), ctx, limit)
}

// Create mocks base method
func (m *MockBeatmapSet) Create(ctx context.Context, from interface{}) (*entity.BeatmapSetFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, from)
	ret0, _ := ret[0].(*entity.BeatmapSetFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockBeatmapSetMockRecorder) Create(ctx, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockBeatmapSet)(nil).Create), ctx, from)
}

// Update mocks base method
func (m *MockBeatmapSet) Update(ctx context.Context, id uint, from interface{}) (*entity.BeatmapSetFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, from)
	ret0, _ := ret[0].(*entity.BeatmapSetFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockBeatmapSetMockRecorder) Update(ctx, id, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBeatmapSet)(nil).Update), ctx, id, from)
}

// Delete mocks base method
func (m *MockBeatmapSet) Delete(ctx context.Context, id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBeatmapSetMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBeatmapSet)(nil).Delete), ctx, id)
}

// FetchFromBancho mocks base method
func (m *MockBeatmapSet) FetchFromBancho(ctx context.Context, id uint) (*entity.BeatmapSetFull, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchFromBancho", ctx, id)
	ret0, _ := ret[0].(*entity.BeatmapSetFull)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchFromBancho indicates an expected call of FetchFromBancho
func (mr *MockBeatmapSetMockRecorder) FetchFromBancho(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchFromBancho", reflect.TypeOf((*MockBeatmapSet)(nil).FetchFromBancho), ctx, id)
}

// MockUser is a mock of User interface
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockUser) Get(ctx context.Context, userId uint, mode string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, userId, mode)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockUserMockRecorder) Get(ctx, userId, mode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUser)(nil).Get), ctx, userId, mode)
}

// GetShort mocks base method
func (m *MockUser) GetShort(ctx context.Context, userId uint, mode string) (*entity.UserShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShort", ctx, userId, mode)
	ret0, _ := ret[0].(*entity.UserShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShort indicates an expected call of GetShort
func (mr *MockUserMockRecorder) GetShort(ctx, userId, mode interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShort", reflect.TypeOf((*MockUser)(nil).GetShort), ctx, userId, mode)
}

// GetByBasic mocks base method
func (m *MockUser) GetByBasic(ctx context.Context, login, pwd string) (*entity.UserShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByBasic", ctx, login, pwd)
	ret0, _ := ret[0].(*entity.UserShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByBasic indicates an expected call of GetByBasic
func (mr *MockUserMockRecorder) GetByBasic(ctx, login, pwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByBasic", reflect.TypeOf((*MockUser)(nil).GetByBasic), ctx, login, pwd)
}

// ComputeFields mocks base method
func (m *MockUser) ComputeFields(ctx context.Context, user entity.User) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ComputeFields", ctx, user)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ComputeFields indicates an expected call of ComputeFields
func (mr *MockUserMockRecorder) ComputeFields(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ComputeFields", reflect.TypeOf((*MockUser)(nil).ComputeFields), ctx, user)
}

// Create mocks base method
func (m *MockUser) Create(ctx context.Context, name, email, pwd string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, name, email, pwd)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUserMockRecorder) Create(ctx, name, email, pwd interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUser)(nil).Create), ctx, name, email, pwd)
}

// Update mocks base method
func (m *MockUser) Update(ctx context.Context, userId uint, from interface{}) (*entity.UserShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, userId, from)
	ret0, _ := ret[0].(*entity.UserShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUserMockRecorder) Update(ctx, userId, from interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUser)(nil).Update), ctx, userId, from)
}

// Activate mocks base method
func (m *MockUser) Activate(ctx context.Context, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Activate", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Activate indicates an expected call of Activate
func (mr *MockUserMockRecorder) Activate(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Activate", reflect.TypeOf((*MockUser)(nil).Activate), ctx, userId)
}

// Deactivate mocks base method
func (m *MockUser) Deactivate(ctx context.Context, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deactivate", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deactivate indicates an expected call of Deactivate
func (mr *MockUserMockRecorder) Deactivate(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deactivate", reflect.TypeOf((*MockUser)(nil).Deactivate), ctx, userId)
}

// Ban mocks base method
func (m *MockUser) Ban(ctx context.Context, userId uint, time time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ban", ctx, userId, time)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ban indicates an expected call of Ban
func (mr *MockUserMockRecorder) Ban(ctx, userId, time interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ban", reflect.TypeOf((*MockUser)(nil).Ban), ctx, userId, time)
}

// UnBan mocks base method
func (m *MockUser) UnBan(ctx context.Context, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnBan", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnBan indicates an expected call of UnBan
func (mr *MockUserMockRecorder) UnBan(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnBan", reflect.TypeOf((*MockUser)(nil).UnBan), ctx, userId)
}

// Mute mocks base method
func (m *MockUser) Mute(ctx context.Context, userId uint, time time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mute", ctx, userId, time)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mute indicates an expected call of Mute
func (mr *MockUserMockRecorder) Mute(ctx, userId, time interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mute", reflect.TypeOf((*MockUser)(nil).Mute), ctx, userId, time)
}

// UnMute mocks base method
func (m *MockUser) UnMute(ctx context.Context, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnMute", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnMute indicates an expected call of UnMute
func (mr *MockUserMockRecorder) UnMute(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnMute", reflect.TypeOf((*MockUser)(nil).UnMute), ctx, userId)
}

// UpdateLastVisit mocks base method
func (m *MockUser) UpdateLastVisit(ctx context.Context, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastVisit", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastVisit indicates an expected call of UpdateLastVisit
func (mr *MockUserMockRecorder) UpdateLastVisit(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastVisit", reflect.TypeOf((*MockUser)(nil).UpdateLastVisit), ctx, userId)
}

// MockFriend is a mock of Friend interface
type MockFriend struct {
	ctrl     *gomock.Controller
	recorder *MockFriendMockRecorder
}

// MockFriendMockRecorder is the mock recorder for MockFriend
type MockFriendMockRecorder struct {
	mock *MockFriend
}

// NewMockFriend creates a new mock instance
func NewMockFriend(ctrl *gomock.Controller) *MockFriend {
	mock := &MockFriend{ctrl: ctrl}
	mock.recorder = &MockFriendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFriend) EXPECT() *MockFriendMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockFriend) Add(ctx context.Context, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockFriendMockRecorder) Add(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockFriend)(nil).Add), ctx, userId)
}

// Remove mocks base method
func (m *MockFriend) Remove(ctx context.Context, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", ctx, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockFriendMockRecorder) Remove(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockFriend)(nil).Remove), ctx, userId)
}

// Check mocks base method
func (m *MockFriend) Check(ctx context.Context, userId, targetId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", ctx, userId, targetId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockFriendMockRecorder) Check(ctx, userId, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockFriend)(nil).Check), ctx, userId, targetId)
}

// GetSubscriptions mocks base method
func (m *MockFriend) GetSubscriptions(ctx context.Context, userId uint) (*[]entity.UserShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubscriptions", ctx, userId)
	ret0, _ := ret[0].(*[]entity.UserShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubscriptions indicates an expected call of GetSubscriptions
func (mr *MockFriendMockRecorder) GetSubscriptions(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubscriptions", reflect.TypeOf((*MockFriend)(nil).GetSubscriptions), ctx, userId)
}

// GetFriends mocks base method
func (m *MockFriend) GetFriends(ctx context.Context, userId uint) (*[]entity.UserShort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriends", ctx, userId)
	ret0, _ := ret[0].(*[]entity.UserShort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriends indicates an expected call of GetFriends
func (mr *MockFriendMockRecorder) GetFriends(ctx, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriends", reflect.TypeOf((*MockFriend)(nil).GetFriends), ctx, userId)
}
