package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/imirjar/rb-auth/docs"
	"github.com/imirjar/rb-auth/internal/models"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Service interface {
	BuildJWTString(context.Context, models.User) (string, error)
	Registrate(context.Context, models.User) error
}

type HTTPServer struct {
	Service Service
	Server  *http.Server
}

// @Title RB_AUTH API
// @Description Simple JWT auth.
// @Version 1.0

// @Contact.email support@redbeaver.ru

// @BasePath /api/v1
// @Host localhost:8080
func New() (*HTTPServer, error) {
	gtw := HTTPServer{}

	router := chi.NewRouter()

	router.Post("/login", gtw.LogIn())   // retrun JWT if ok
	router.Post("/signin", gtw.SignIn()) // add user in system

	router.Route("/token", func(token chi.Router) {
		token.Post("/refresh", gtw.Refresh()) // return new jwt with new lifetime
	})

	router.Route("/swagger", func(swagger chi.Router) {
		swagger.Get("/*", httpSwagger.WrapHandler) // return new jwt with new lifetime
	})

	gtw.Server = &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	return &gtw, nil
}

// HelloHandler пример обработчика
// @Summary Registrate new user
// @Description Add new user login and password
// @Parameters sdf
// @Tags JWT
// @Success 200 {string} string "success"
// @Router /signin [post]
func (s *HTTPServer) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User
		// Parse r.Bidy to models.User struct. User must be valid!
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, errIncorrectUser.Error(), http.StatusBadRequest)
			return
		}

		// Trying to add new user into system
		err = s.Service.Registrate(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}
}

// GoodbyeHandler пример обработчика
// @Summary Get user JWT
// @Description Authentificate user by login and password and retrun JWT if ok
// @Tags JWT
// @Success 200 {string} string "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjY1NzI1NjksIlVzZXJJRCI6MX0.GAUD2ulqg-UIXsomcc6B9vFD5Eqyrg75jwjH39o4BXg"
// @Router /login [post]
func (s *HTTPServer) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User

		// Parse r.Body to models.User struct. User must be valid!
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := s.Service.BuildJWTString(r.Context(), user)
		// if user isn't exist
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// response
		if err = json.NewEncoder(w).Encode(token); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

// Token:

// Return new jwt with new lifetime
func (s *HTTPServer) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
