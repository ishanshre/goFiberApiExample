package repository

import "github.com/ishanshre/goFiberApiExample/internals/models"

type DatabaseRepo interface {
	GetAllProducts(limit, offset int) ([]*models.Product, error)
	InsertProduct(p *models.Product) error
	DeleteProduct(*models.Product) error
	UpdateProduct(*models.Product) error
	GetProductByID(int) (*models.Product, error)
}
