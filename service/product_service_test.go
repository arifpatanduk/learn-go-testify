package service

import (
	"golang-unit-testing/entity"
	"golang-unit-testing/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductService{Repository: productRepository}

func TestProductServiceGetOneProductNotFound(t *testing.T) {
	productRepository.Mock.On("FindById", "1").Return(nil)

	product, err := productService.GetOneProduct("1")

	assert.Nil(t, product)
	assert.NotNil(t, err)
	assert.Equal(t, "product not found", err.Error(), "error response has to be 'product not found'")
}

func TestProductServiceGetOneProduct(t *testing.T) {
	product := entity.Product{
		Id: "2",
		Name: "Kaca mata",
	}

	productRepository.Mock.On("FindById", "2").Return(product)

	result, err := productService.GetOneProduct("2")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, product.Id, result.Id, "result has to be 2")
}

func TestProductServiceGetAllProduct(t *testing.T) {
	products := []entity.Product{
		{
			Id: "1", Name: "Masker",
		},
		{
			Id: "2",
			Name: "Kaca mata",
		},
	}

	productRepository.Mock.On("FindAll").Return(products)

	result, _ := productService.GetAllProduct()

	assert.NotNil(t, result)
	assert.Equal(t, len(products), len(result), "Product has to be 2")
	assert.Equal(t, products, result, "Product has to be retrieved successfully")

	for i, v := range result {
		assert.Equal(t, products[i].Id, v.Id, "same id")
		assert.Equal(t, products[i].Name, v.Name, "same name")
	}

}

func TestProductServiceGetAllProductFail(t *testing.T) {
	products := []entity.Product{
		{
			Id: "1", Name: "Masker",
		},
		{
			Id: "2",
			Name: "Kaca mata",
		},
	}

	productRepository.Mock.On("FindAll").Return(products)

	result, _ := productService.GetAllProduct()

	assert.NotEqual(t, len(products), len(result) + 1, "gagal")
	
	for i, v := range result {
		assert.NotEqual(t, products[i].Name, v.Name + "a", "gagal")
	}
}