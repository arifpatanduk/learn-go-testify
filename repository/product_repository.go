package repository

import "golang-unit-testing/entity"

type ProductRepository interface {
	FindById(id string) *entity.Product
	FindAll() []entity.Product
}