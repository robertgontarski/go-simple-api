package server

import (
	"github.com/go-playground/validator/v10"
	"simple-api/storage"
)

type Server interface {
	Listen() error
}

type ConfigServer struct {
	Addr     string
	Db       storage.Storage
	Validate *validator.Validate
}

func NewConfigServer(addr string, store storage.Storage, validate *validator.Validate) *ConfigServer {
	return &ConfigServer{
		Addr:     addr,
		Db:       store,
		Validate: validate,
	}
}
