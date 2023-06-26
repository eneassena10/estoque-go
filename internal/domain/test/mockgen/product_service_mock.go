// Code generated by MockGen. DO NOT EDIT.
// Source: ./service.go

// Package mockgen is a generated GoMock package.
package mockgen

import (
	reflect "reflect"

	entities "github.com/eneassena10/estoque-go/internal/domain/product/entities"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockIPoductService is a mock of IPoductService interface.
type MockIPoductService struct {
	ctrl     *gomock.Controller
	recorder *MockIPoductServiceMockRecorder
}

// MockIPoductServiceMockRecorder is the mock recorder for MockIPoductService.
type MockIPoductServiceMockRecorder struct {
	mock *MockIPoductService
}

// NewMockIPoductService creates a new mock instance.
func NewMockIPoductService(ctrl *gomock.Controller) *MockIPoductService {
	mock := &MockIPoductService{ctrl: ctrl}
	mock.recorder = &MockIPoductServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPoductService) EXPECT() *MockIPoductServiceMockRecorder {
	return m.recorder
}

// CreateProducts mocks base method.
func (m *MockIPoductService) CreateProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProducts", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProducts indicates an expected call of CreateProducts.
func (mr *MockIPoductServiceMockRecorder) CreateProducts(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProducts", reflect.TypeOf((*MockIPoductService)(nil).CreateProducts), ctx, product)
}

// DeleteProducts mocks base method.
func (m *MockIPoductService) DeleteProducts(ctx *gin.Context, product *entities.ProductRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProducts", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProducts indicates an expected call of DeleteProducts.
func (mr *MockIPoductServiceMockRecorder) DeleteProducts(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProducts", reflect.TypeOf((*MockIPoductService)(nil).DeleteProducts), ctx, product)
}

// GetProductsAll mocks base method.
func (m *MockIPoductService) GetProductsAll(ctx *gin.Context) *[]entities.ProductRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsAll", ctx)
	ret0, _ := ret[0].(*[]entities.ProductRequest)
	return ret0
}

// GetProductsAll indicates an expected call of GetProductsAll.
func (mr *MockIPoductServiceMockRecorder) GetProductsAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsAll", reflect.TypeOf((*MockIPoductService)(nil).GetProductsAll), ctx)
}

// GetProductsOne mocks base method.
func (m *MockIPoductService) GetProductsOne(ctx *gin.Context, product *entities.ProductRequest) *entities.ProductRequest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductsOne", ctx, product)
	ret0, _ := ret[0].(*entities.ProductRequest)
	return ret0
}

// GetProductsOne indicates an expected call of GetProductsOne.
func (mr *MockIPoductServiceMockRecorder) GetProductsOne(ctx, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsOne", reflect.TypeOf((*MockIPoductService)(nil).GetProductsOne), ctx, product)
}

// UpdateProductsCount mocks base method.
func (m *MockIPoductService) UpdateProductsCount(ctx *gin.Context, oldProduct *entities.ProductRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProductsCount", ctx, oldProduct)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProductsCount indicates an expected call of UpdateProductsCount.
func (mr *MockIPoductServiceMockRecorder) UpdateProductsCount(ctx, oldProduct interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProductsCount", reflect.TypeOf((*MockIPoductService)(nil).UpdateProductsCount), ctx, oldProduct)
}