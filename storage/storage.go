package storage

import "simple-api/models"

type Storage interface {
	AddProduct(*models.Product) error
	GetProductByID(int) (*models.Product, error)
	AddClient(*models.Client) error
	GetClientByID(int) (*models.Client, error)
	GetClientByEmail(string) (*models.Client, error)
}
