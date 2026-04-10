package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/imirjar/rb-auth/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type HttpServer struct {
	Service Service
	Port    string
}

// @Title RB_AUTH API
// @Description Simple JWT auth.
// @Version 1.0

// @license.name  GNU GPL 3.0
// @license.url   https://www.gnu.org/licenses/gpl-3.0.html#license-text

// @Contact.email support@redbeaver.ru

// @BasePath /
// @Host localhost:8080

func New(ctx context.Context, port string) *HttpServer {
	return &HttpServer{
		Port: port,
	}
}
func (srv *HttpServer) Run(ctx context.Context) error {

	router := mux.NewRouter()

	// Auth handlers
	auth := router.PathPrefix("/auth").Subrouter()
	auth.Handle("/validate", srv.authHandler(ctx)).Methods("POST")

	user := router.PathPrefix("/users").Subrouter()
	user.Handle("/{id}", srv.userHandler(ctx)).Methods("POST")
	user.Handle("/", srv.usersHandler(ctx)).Methods("POST", "GET")

	// RB_AUTH API SWAGGER
	router.Handle("/swagger/", httpSwagger.Handler()).Methods("GET")
	router.Handle("/health", srv.healthHandler(ctx)).Methods("GET")

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", srv.Port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// Wait for shutdown signal
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(shutdownCtx)
}

func (srv *HttpServer) healthHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
