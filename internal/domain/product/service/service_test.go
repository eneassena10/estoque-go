package service

import (
	"errors"
	"testing"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	"github.com/eneassena10/estoque-go/internal/domain/test/mockgen"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewProductService(t *testing.T) {
	productMocked := &entities.ProductRequest{
		ID:         1,
		Name:       "Sapato",
		Price:      600,
		Quantidade: 1,
	}

	t.Run("GetProductsAll", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		productsListMocked := new([]entities.ProductRequest)
		*productsListMocked = append(*productsListMocked, *productMocked)

		productRepositoryMocked.EXPECT().GetProductsAll(gomock.Any()).Return(productsListMocked)

		serviceMocked := NewProductService(productRepositoryMocked)
		prod := serviceMocked.GetProductsAll(nil)
		assert.NotNil(t, prod)
		assert.True(t, len(*prod) == 1)
	})
	t.Run("GetProductsOne", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		productSearchMocked := &entities.ProductRequest{ID: 1}

		productRepositoryMocked.EXPECT().GetProductsOne(gomock.Any(), productSearchMocked).Return(productMocked)

		serviceMocked := NewProductService(productRepositoryMocked)
		productResult := serviceMocked.GetProductsOne(nil, productSearchMocked)
		assert.NotNil(t, productResult)
		assert.Equal(t, productSearchMocked.ID, productMocked.ID)
	})
	t.Run("CreateProducts - success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		productNewMocked := &entities.ProductRequest{
			Name:       "Tenis",
			Price:      450,
			Quantidade: 10,
		}

		productRepositoryMocked.EXPECT().CreateProducts(gomock.Any(), productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.CreateProducts(nil, productNewMocked)
		assert.Nil(t, erro)
		assert.NoError(t, erro)
	})
	t.Run("CreateProducts - Fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		errorExpected := errors.New("não possível criar um novo produto! 'Tente Novamente'")
		productNewMocked := &entities.ProductRequest{
			Name:  "Tenis",
			Price: 450,
		}

		productRepositoryMocked.EXPECT().CreateProducts(gomock.Any(), productNewMocked).Return(errorExpected)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.CreateProducts(nil, productNewMocked)
		assert.Error(t, erro)
	})
	t.Run("UpdateProductsCount", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		oldProductMocked := &entities.ProductRequest{
			ID:         1,
			Quantidade: 5,
			Name:       "Tenis",
			Price:      450,
		}
		productNewMocked := &entities.ProductRequest{
			ID:         1,
			Quantidade: 5,
		}

		productRepositoryMocked.EXPECT().GetProductsOne(gomock.Any(), productNewMocked).Return(oldProductMocked)
		productRepositoryMocked.EXPECT().UpdateProductsCount(gomock.Any(), oldProductMocked, productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.UpdateProductsCount(nil, productNewMocked)
		assert.NoError(t, erro)
	})
	t.Run("UpdateProductsCount or fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		context := &gin.Context{}

		productNewMocked := &entities.ProductRequest{
			ID:         1,
			Quantidade: 5,
		}
		expectError := errors.New("not found product")
		productRepositoryMocked.EXPECT().GetProductsOne(context, productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.UpdateProductsCount(context, productNewMocked)
		assert.Error(t, erro)
		assert.EqualValues(t, expectError.Error(), erro.Error())
	})
	t.Run("DeleteProducts", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		context := &gin.Context{}

		productNewMocked := &entities.ProductRequest{
			ID:         1,
			Quantidade: 5,
		}
		productRepositoryMocked.EXPECT().DeleteProducts(context, productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.DeleteProducts(context, productNewMocked)
		assert.NoError(t, erro)
	})
}
