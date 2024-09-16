package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/imirjar/rb-auth/internal/entity/models"
)

type Service interface {
	Authenticate(context.Context, models.User) (models.User, error)
	Authorize(context.Context, models.User)
	Registrate(context.Context, models.User) error
}

type HTTPServer struct {
	Service Service
	Server  *http.Server
}

func New() (*HTTPServer, error) {
	gtw := HTTPServer{}

	router := chi.NewRouter()
	router.Post("/auth", gtw.LogIn())
	router.Post("/register", gtw.SignIn())

	// router.Route("/token", func(token chi.Router) {
	// 	token.Post("/refresh", gtw.Refresh())
	// })

	gtw.Server = &http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	return &gtw, nil
}

func (s *HTTPServer) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User

		// Parse r.Bidy to models.User struct. User must be valid!
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, errIncorrectUser.Error(), http.StatusBadRequest)
			return
		} else if !user.IsValid() {
			http.Error(w, errInvalidUser.Error(), http.StatusBadRequest)
			return
		}

		log.Println(user)

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

func (s *HTTPServer) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user models.User

		// Parse r.Bidy to models.User struct. User must be valid!
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		} else if !user.IsValid() {
			http.Error(w, "user is not valid", http.StatusBadRequest)
			return
		}

		log.Println(user)

		authUser, err := s.Service.Authenticate(r.Context(), user)
		// if user isn't exist
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// response
		if err = json.NewEncoder(w).Encode(authUser); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

// func (s *HTTPServer) Refresh() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var user models.User

// 		err := json.NewDecoder(r.Body).Decode(&user)
// 		if err != nil {
// 			log.Printf("Decode error: %s", err)
// 		}

// 		s.Service.Authorize(r.Context(), user)
// 	}
// }
