package service

import (
	"errors"
	"testing"

	"github.com/eneassena10/estoque-go/internal/domain/product/entities"
	"github.com/eneassena10/estoque-go/internal/domain/test/mockgen"
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

		productRepositoryMocked.EXPECT().GetProductsAll().Return(productsListMocked)

		serviceMocked := NewProductService(productRepositoryMocked)
		prod := serviceMocked.GetProductsAll()
		assert.NotNil(t, prod)
		assert.True(t, len(*prod) == 1)
	})
	t.Run("GetProductsOne", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)
		productSearchMocked := &entities.ProductRequest{ID: 1}

		productRepositoryMocked.EXPECT().GetProductsOne(productSearchMocked).Return(productMocked)

		serviceMocked := NewProductService(productRepositoryMocked)
		productResult := serviceMocked.GetProductsOne(productSearchMocked)
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

		productRepositoryMocked.EXPECT().CreateProducts(productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.CreateProducts(productNewMocked)
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

		productRepositoryMocked.EXPECT().CreateProducts(productNewMocked).Return(errorExpected)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.CreateProducts(productNewMocked)
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

		productRepositoryMocked.EXPECT().GetProductsOne(productNewMocked).Return(oldProductMocked)
		productRepositoryMocked.EXPECT().UpdateProductsCount(oldProductMocked, productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.UpdateProductsCount(productNewMocked)
		assert.NoError(t, erro)
	})
	t.Run("UpdateProductsCount or fail", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)

		productNewMocked := &entities.ProductRequest{
			ID:         1,
			Quantidade: 5,
		}
		expectError := errors.New("not found product")
		productRepositoryMocked.EXPECT().GetProductsOne(productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.UpdateProductsCount(productNewMocked)
		assert.Error(t, erro)
		assert.EqualValues(t, expectError.Error(), erro.Error())
	})
	t.Run("DeleteProducts", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		productRepositoryMocked := mockgen.NewMockIProductRepository(ctrl)

		productNewMocked := &entities.ProductRequest{
			ID:         1,
			Quantidade: 5,
		}
		productRepositoryMocked.EXPECT().DeleteProducts(productNewMocked).Return(nil)

		serviceMocked := NewProductService(productRepositoryMocked)
		erro := serviceMocked.DeleteProducts(productNewMocked)
		assert.NoError(t, erro)
	})
}
