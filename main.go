package main

import (
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

	api := server.NewApiServer(
		server.NewConfigServer(
			os.Getenv("SERVER_LISTEN"),
			store,
		),
	)

	if err := api.Listen(); err != nil {
		slog.Error("error while listening", "err", err)
		return
	}
}
