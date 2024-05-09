package server

import (
	"net/http"
	"simple-api/handler"
)

type ApiServer struct {
	*ConfigServer
	router *http.ServeMux
}

func NewApiServer(conf *ConfigServer) *ApiServer {
	return &ApiServer{
		ConfigServer: conf,
		router:       http.NewServeMux(),
	}
}

func (s *ApiServer) Listen() error {
	srv := http.Server{
		Addr:    s.Addr,
		Handler: s.router,
	}

	s.setupRouter()

	return srv.ListenAndServe()
}

func (s *ApiServer) setupRouter() {
	productHandler := handler.NewProductHandler(s.Db, s.Validate)

	s.router.HandleFunc("POST /product", handler.AsJSONContent(handler.MakeHandleFunc(productHandler.CreateProductHandler)))
	s.router.HandleFunc("GET /product/{id}", handler.AsJSONContent(handler.MakeHandleFunc(productHandler.GetProductByIDHandler)))
}
