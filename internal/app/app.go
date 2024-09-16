package app

import (
	"log"

	gateway "github.com/imirjar/rb-auth/internal/gateway/http"
	service "github.com/imirjar/rb-auth/internal/service/jwt"
	storage "github.com/imirjar/rb-auth/internal/storage/memory"
)

func Run() error {

	storage, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}

	service, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	gw, err := gateway.New()
	if err != nil {
		log.Fatal(err)
	}

	service.Storage = storage
	gw.Service = service
	return gw.Server.ListenAndServe()
}
