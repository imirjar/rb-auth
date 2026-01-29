package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/imirjar/rb-auth/internal/domain/entities"
)

type Service interface {
	LogIn(context.Context, entities.LoginRequest) (entities.TokenPair, error)

	Validate(context.Context, string) (bool, error)
	Refresh(context.Context, string) (entities.TokenPair, error)
}

// @Tags JWT
// @Router /login [post]
// @Summary Get user JWT
// @Description Authentificate user by login and password and retrun JWT if ok
// @Param user body models.User true "query params"
// @Success 200 {string} string "success"
// @Failure 400  {string}  string    "user isn't correct"
// @Failure 403  {string}  string    "user isn't valid"
// @Failure 500  {string}  string    "some error"
func (s *HTTPServer) LogIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var login entities.LoginRequest

		// Parse r.Body to models.User struct. User must be valid!
		err := json.NewDecoder(r.Body).Decode(&login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := s.Service.LogIn(r.Context(), login)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (s *HTTPServer) SignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var signin entities.SigninRequest

		// Parse r.Body to models.User struct. User must be valid!
		err := json.NewDecoder(r.Body).Decode(&signin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = s.Service.SignIn(r.Context(), signin)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// // response
		// w.Header().Set("Content-Type", "application/json")
		// if err = json.NewEncoder(w).Encode(resp); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		if err := json.NewEncoder(w).Encode(""); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

// @Tags JWT
// @Router /token/validate [post]
// @Summary Validate user JWT
// @Description Authentificate user by login and password and retrun JWT if ok
// @Param user body models.User true "query params"
// @Success 200 {string} string "Token is vali"
// @Failure 401  {string}  string    "Missing Authorization Header"
// @Failure 403  {string}  string    "user isn't valid"
// @Failure 500  {string}  string    "some error"
func (s *HTTPServer) Validate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var token string

		err := json.NewDecoder(r.Body).Decode(&token)
		if err != nil {
			http.Error(w, errIncorrectUser.Error(), http.StatusBadRequest)
			return
		}

		user, err := s.Service.Validate(context.Background(), token)
		if err != nil {
			http.Error(w, errIncorrectUser.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// @Tags JWT
// @Router /token/refresh [post]
// @Summary Refresh user JWT
// @Description Send your JWT to prolongate your JWT expired period
// @Param user body models.User true "query params"
// @Success 200 {string} string "success"
// @Failure 400  {string}  string    "user isn't correct"
// @Failure 403  {string}  string    "user isn't valid"
// @Failure 500  {string}  string    "some error"
func (s *HTTPServer) Refresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token entities.TokenPair

		err := json.NewDecoder(r.Body).Decode(&token)
		if err != nil {
			http.Error(w, errIncorrectUser.Error(), http.StatusBadRequest)
			return
		}

		newToken, err := s.Service.Refresh(context.Background(), token.Refresh)
		if err != nil {
			http.Error(w, errIncorrectUser.Error(), http.StatusBadRequest)
			return
		}

		if err := json.NewEncoder(w).Encode(newToken); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
