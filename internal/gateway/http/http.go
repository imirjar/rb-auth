package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/rb-auth/internal/entity/models"
)

type service interface {
	Authenticate(context.Context, models.User)
	Authorize(context.Context, models.User)
	Registrate(context.Context, models.User)
}

type HTTPServer struct {
	Service service
	Server  *http.Server
}

func New() (*HTTPServer, error) {
	gtw := HTTPServer{}

	router := chi.NewRouter()
	router.Route("/auth", func(auth chi.Router) {
		auth.Post("/", gtw.Authenticate())
	})
	router.Route("/token", func(token chi.Router) {
		token.Post("/", gtw.Authorize())
	})
	router.Route("/register", func(register chi.Router) {
		register.Post("/", gtw.Registrate())
	})

	gtw.Server = &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	return &gtw, nil
}

func (s *HTTPServer) Registrate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Decode error: %s", err)
		}

		s.Service.Registrate(r.Context(), user)
	}
}

func (s *HTTPServer) Authenticate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Decode error: %s", err)
		}

		s.Service.Authenticate(r.Context(), user)
	}
}

func (s *HTTPServer) Authorize() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.Printf("Decode error: %s", err)
		}

		s.Service.Authorize(r.Context(), user)
	}
}
