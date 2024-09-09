package app

import (
	"log"

	gateway "github.com/imirjar/rb-auth/internal/gateway/http"
	"github.com/imirjar/rb-auth/internal/service"
)

func Run() error {

	service, err := service.New()
	if err != nil {
		log.Fatal(err)
	}

	gw, err := gateway.New()
	if err != nil {
		log.Fatal(err)
	}

	gw.Service = service
	return gw.Server.ListenAndServe()
}
