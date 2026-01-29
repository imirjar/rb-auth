package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/imirjar/rb-auth/docs"
	"github.com/imirjar/rb-auth/internal/service"
	httpSwagger "github.com/swaggo/http-swagger"
)

type HTTPServer struct {
	Service service.Service
	Server  *http.Server
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

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Auth handlers
	router.Post("/login", gtw.LogIn())   // retrun JWT if ok and create session
	router.Post("/signin", gtw.SignIn()) // retrun JWT if ok and create user

	// Manipulations with JWT
	router.Route("/token", func(token chi.Router) {
		token.Post("/refresh", gtw.Refresh())
		token.Post("/validate", gtw.Validate()) // return true if jwt if valid
	})

	// RB_AUTH API SWAGGER
	router.Route("/api", func(swagger chi.Router) {
		swagger.Get("/v1/*", httpSwagger.WrapHandler)
	})

	gtw.Server = &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	fmt.Printf("App run on port %s", gtw.Server.Addr)

	return &gtw, nil
}
