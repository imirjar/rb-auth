package app

import (
	"log"

	"github.com/imirjar/rb-auth/config"
	gateway "github.com/imirjar/rb-auth/internal/gateway/http"
	tokenService "github.com/imirjar/rb-auth/internal/service/token"
	userService "github.com/imirjar/rb-auth/internal/service/user"
	storage "github.com/imirjar/rb-auth/internal/storage/memory"
)

func Run() error {
	config := config.New()
	// log.Print(config)

	storage, err := storage.New()
	if err != nil {
		log.Fatal(err)
	}

	userService, err := userService.New()
	if err != nil {
		log.Fatal(err)
	}

	// log.Print(config.Security.Pub.Key)
	tokenService, err := tokenService.New(config.Security.Priv.Key, config.Security.Pub.Key)
	if err != nil {
		log.Fatal(err)
	}

	gw, err := gateway.New(config.Http.Port)
	if err != nil {
		log.Fatal(err)
	}

	userService.Storage = storage
	// tokenService.Storage = storage
	gw.UserService = userService
	gw.TokenService = tokenService
	return gw.Server.ListenAndServe()
}
