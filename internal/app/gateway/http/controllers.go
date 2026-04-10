package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/imirjar/rb-auth/internal/domain"
)

type Service interface {
	CreateUser(context.Context, domain.User) (domain.User, error)
	ReadUser()
	ReadUsers()
	UpdateUser(context.Context, domain.User) (domain.User, error)
	DeleteUser(context.Context, string) (domain.User, error)

	VerifyUser(context.Context, domain.User) (bool, error)
}

func (srv *HttpServer) authHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var user domain.User
			var err error

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "GET":
			var users []domain.User
			var err error

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func (srv *HttpServer) userHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var user domain.User
			var err error

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(user); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "GET":
			var users []domain.User
			var err error

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}

func (srv *HttpServer) usersHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			var user domain.User
			var err error

			if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
				http.Error(w, "Invalid JSON format: "+err.Error(), http.StatusBadRequest)
				return
			}
			defer r.Body.Close()

			createUser, err := srv.Service.CreateUser(ctx, user)

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(createUser); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		case "GET":
			var users []domain.User
			var err error

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			if err = json.NewEncoder(w).Encode(users); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
