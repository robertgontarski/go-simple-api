package server

import (
	"github.com/go-playground/validator/v10"
	"net/http"
	"simple-api/auth"
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
	JWTAuth := auth.NewJWTAuth()
	validate := validator.New(validator.WithRequiredStructEnabled())
	dh := handler.NewDefaultHandler(s.Db, validate, JWTAuth)

	authHandler := handler.NewAuthHandler(dh)

	s.router.HandleFunc("POST /login", handler.AsJSONContent(handler.MakeHandleFunc(authHandler.LoginHandler)))

	productHandler := handler.NewProductHandler(dh)

	s.router.HandleFunc("POST /product", handler.AsJSONContent(handler.MakeHandleFunc(productHandler.CreateProductHandler)))
	s.router.HandleFunc("GET /product/{id}", handler.AsProtected(handler.AsJSONContent(handler.MakeHandleFunc(productHandler.GetProductByIDHandler)), s.Db))

	clientHandler := handler.NewClientHandler(dh)

	s.router.HandleFunc("POST /client", handler.AsJSONContent(handler.MakeHandleFunc(clientHandler.CreateClientHandler)))
	s.router.HandleFunc("GET /client/{id}", handler.AsJSONContent(handler.MakeHandleFunc(clientHandler.GetClientByIDHandler)))
}
