package storage

import "simple-api/models"

type Storage interface {
	AddProduct(*models.Product) error
	//UpdateProduct(*models.Product) error
	//DeleteProduct(int) error
	GetProductByID(int) (*models.Product, error)
	//GetProducts() ([]*models.Product, error)
}
