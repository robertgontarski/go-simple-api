package handler

import (
	"github.com/go-playground/validator/v10"
	"simple-api/auth"
	"simple-api/storage"
)

type DefaultHandler struct {
	db       storage.Storage
	validate *validator.Validate
	jwtAuth  *auth.JWTAuth
}

func NewDefaultHandler(db storage.Storage, validate *validator.Validate, jwtAuth *auth.JWTAuth) *DefaultHandler {
	return &DefaultHandler{
		db:       db,
		validate: validate,
		jwtAuth:  jwtAuth,
	}
}
