// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go
//
// Generated by this command:
//
//	mockgen -source=auth.go -destination=mock/auth.go -package=mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	domain "github.com/bagashiz/go-pos/internal/core/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockTokenService is a mock of TokenService interface.
type MockTokenService struct {
	ctrl     *gomock.Controller
	recorder *MockTokenServiceMockRecorder
}

// MockTokenServiceMockRecorder is the mock recorder for MockTokenService.
type MockTokenServiceMockRecorder struct {
	mock *MockTokenService
}

// NewMockTokenService creates a new mock instance.
func NewMockTokenService(ctrl *gomock.Controller) *MockTokenService {
	mock := &MockTokenService{ctrl: ctrl}
	mock.recorder = &MockTokenServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTokenService) EXPECT() *MockTokenServiceMockRecorder {
	return m.recorder
}

// CreateToken mocks base method.
func (m *MockTokenService) CreateToken(user *domain.User) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateToken", user)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateToken indicates an expected call of CreateToken.
func (mr *MockTokenServiceMockRecorder) CreateToken(user any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateToken", reflect.TypeOf((*MockTokenService)(nil).CreateToken), user)
}

// VerifyToken mocks base method.
func (m *MockTokenService) VerifyToken(token string) (*domain.TokenPayload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", token)
	ret0, _ := ret[0].(*domain.TokenPayload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockTokenServiceMockRecorder) VerifyToken(token any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockTokenService)(nil).VerifyToken), token)
}

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockAuthService) Login(ctx context.Context, email, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, email, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthServiceMockRecorder) Login(ctx, email, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuthService)(nil).Login), ctx, email, password)
}
