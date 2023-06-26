// Code generated by MockGen. DO NOT EDIT.
// Source: ./controllers_products.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIProductControllers is a mock of IProductControllers interface.
type MockIProductControllers struct {
	ctrl     *gomock.Controller
	recorder *MockIProductControllersMockRecorder
}

// MockIProductControllersMockRecorder is the mock recorder for MockIProductControllers.
type MockIProductControllersMockRecorder struct {
	mock *MockIProductControllers
}

// NewMockIProductControllers creates a new mock instance.
func NewMockIProductControllers(ctrl *gomock.Controller) *MockIProductControllers {
	mock := &MockIProductControllers{ctrl: ctrl}
	mock.recorder = &MockIProductControllersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProductControllers) EXPECT() *MockIProductControllersMockRecorder {
	return m.recorder
}

// CreateProducts mocks base method.
func (m *MockIProductControllers) CreateProducts(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateProducts", ctx)
}

// CreateProducts indicates an expected call of CreateProducts.
func (mr *MockIProductControllersMockRecorder) CreateProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProducts", reflect.TypeOf((*MockIProductControllers)(nil).CreateProducts), ctx)
}

// DeleteProducts mocks base method.
func (m *MockIProductControllers) DeleteProducts(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteProducts", ctx)
}

// DeleteProducts indicates an expected call of DeleteProducts.
func (mr *MockIProductControllersMockRecorder) DeleteProducts(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProducts", reflect.TypeOf((*MockIProductControllers)(nil).DeleteProducts), ctx)
}

// GetProductsAll mocks base method.
func (m *MockIProductControllers) GetProductsAll(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetProductsAll", ctx)
}

// GetProductsAll indicates an expected call of GetProductsAll.
func (mr *MockIProductControllersMockRecorder) GetProductsAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsAll", reflect.TypeOf((*MockIProductControllers)(nil).GetProductsAll), ctx)
}

// GetProductsByID mocks base method.
func (m *MockIProductControllers) GetProductsByID(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetProductsByID", ctx)
}

// GetProductsByID indicates an expected call of GetProductsByID.
func (mr *MockIProductControllersMockRecorder) GetProductsByID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsByID", reflect.TypeOf((*MockIProductControllers)(nil).GetProductsByID), ctx)
}

// UpdateProductsCount mocks base method.
func (m *MockIProductControllers) UpdateProductsCount(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateProductsCount", ctx)
}

// UpdateProductsCount indicates an expected call of UpdateProductsCount.
func (mr *MockIProductControllersMockRecorder) UpdateProductsCount(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductsCount", reflect.TypeOf((*MockIProductControllers)(nil).UpdateProductsCount), ctx)
}