package server

import (
	"simple-api/storage"
)

type Server interface {
	Listen() error
}

type ConfigServer struct {
	Addr string
	Db   storage.Storage
}

func NewConfigServer(addr string, store storage.Storage) *ConfigServer {
	return &ConfigServer{
		Addr: addr,
		Db:   store,
	}
}
