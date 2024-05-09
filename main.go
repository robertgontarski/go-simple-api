package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"simple-api/server"
	"simple-api/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("error while load env", "err", err)
		return
	}

	store, err := storage.NewMysqlStore()
	if err != nil {
		slog.Error("error while connecting into db", "err", err)
		return
	}

	validate := validator.New(validator.WithRequiredStructEnabled())

	api := server.NewApiServer(
		server.NewConfigServer(
			os.Getenv("SERVER_LISTEN"),
			store,
			validate,
		),
	)

	if err := api.Listen(); err != nil {
		slog.Error("error while listening", "err", err)
		return
	}
}
