package handler

import (
	"github.com/go-playground/validator/v10"
	"simple-api/storage"
)

type DefaultHandler struct {
	db       storage.Storage
	validate *validator.Validate
}

func NewDefaultHandler(db storage.Storage, validate *validator.Validate) *DefaultHandler {
	return &DefaultHandler{
		db:       db,
		validate: validate,
	}
}
