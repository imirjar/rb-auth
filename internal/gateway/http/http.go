package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	_ "github.com/imirjar/rb-auth/docs"
	"github.com/imirjar/rb-auth/internal/models"
	httpSwagger "github.com/swaggo/http-swagger"
)

type UserService interface {
	CheckUser(context.Context, models.User) (models.User, bool) //CheckUser
	AddUser(context.Context, models.User) error                 //AddUser
}

type TokenService interface {
	Create(context.Context, models.User) (string, error)
	Refresh(context.Context, string) (string, error)
	Read(context.Context, string) (models.User, error)
	Validate(context.Context, string) bool
}

type SessionService interface {
	CreateSession(context.Context, models.User, string)
	DeleteSession(context.Context, string)
}

type HTTPServer struct {
	UserService    UserService
	TokenService   TokenService
	SessionService SessionService
	Server         *http.Server
}

// @Title RB_AUTH API
// @Description Simple JWT auth.
// @Version 1.0

// @license.name  GNU GPL 3.0
// @license.url   https://www.gnu.org/licenses/gpl-3.0.html#license-text

// @Contact.email support@redbeaver.ru

// @BasePath /
// @Host localhost:8080
func New(port string) (*HTTPServer, error) {
	gtw := HTTPServer{}

	router := chi.NewRouter()

	// Auth handlers
	router.Post("/login", gtw.LogIn())   // retrun JWT if ok and create session
	router.Post("/signin", gtw.SignIn()) // add user in system
	router.Post("/logout", gtw.SignIn()) // delete session

	// Manipulations with JWT
	router.Route("/token", func(token chi.Router) {
		token.Post("/refresh", gtw.Refresh())   // return new jwt with new lifetime
		token.Post("/validate", gtw.Validate()) // return new jwt with new lifetime
	})

	// RB_AUTH API SWAGGER
	router.Route("/api", func(swagger chi.Router) {
		swagger.Get("/v1/*", httpSwagger.WrapHandler) // return new jwt with new lifetime
	})

	gtw.Server = &http.Server{
		Handler: router,
		Addr:    port,
	}

	fmt.Printf("App run on port %s", port)

	return &gtw, nil
}

// @Tags JWT
// @Router /signin [post]
// @Summary Registrate new user
// @Description Add new user login and password
// @Param user body models.User true "query params"
// @Success 200 {string} string "success"
// @Failure 400  {string}  string    "user isn't correct"
// @Failure 500  {string}  string    "some error"
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
		err = s.UserService.AddUser(r.Context(), user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("success"))
	}
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

		var user models.User

		// Parse r.Body to models.User struct. User must be valid!
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		user, ok := s.UserService.CheckUser(r.Context(), user)
		if !ok {
			http.Error(w, errUserIsNotExists.Error(), http.StatusForbidden)
			return
		}

		token, err := s.TokenService.Create(r.Context(), user)
		// if user isn't exist
		if err != nil {
			http.Error(w, errInvalidUser.Error(), http.StatusForbidden)
			return
		}

		// CREATE SESSION

		// response
		if err = json.NewEncoder(w).Encode(token); err != nil {
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
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, errMissingHeader.Error(), http.StatusUnauthorized)
			return
		}

		// expected - "Bearer {token}"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		isValid := s.TokenService.Validate(context.Background(), tokenString)
		if !isValid {
			http.Error(w, errInvalidToken.Error(), http.StatusForbidden)
			return
		}

		prolongJWT, err := s.TokenService.Refresh(r.Context(), tokenString)
		if authHeader == "" {
			log.Println(err)
			http.Error(w, errInternal.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(prolongJWT))

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

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, errMissingHeader.Error(), http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		user, err := s.TokenService.Read(context.Background(), tokenString)
		if err != nil {
			http.Error(w, errInvalidToken.Error(), http.StatusForbidden)
			return
		}

		if err := json.NewEncoder(w).Encode(user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
