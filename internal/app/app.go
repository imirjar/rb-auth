package app

import (
	"log"

	"github.com/imirjar/rb-auth/config"
	gateway "github.com/imirjar/rb-auth/internal/app/gateway/http"
	service "github.com/imirjar/rb-auth/internal/service"
	storage "github.com/imirjar/rb-auth/internal/storage"
)

func Run() error {
	config := config.New()
	// log.Print(config)

	storage, err := storage.New(config.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	service, err := service.New(config.Secret)
	if err != nil {
		log.Fatal(err)
	}

	gw, err := gateway.New(config.Port)
	if err != nil {
		log.Fatal(err)
	}

	service.Storage = storage
	gw.Service = *service
	return gw.Server.ListenAndServe()
}
