package service

import (
	"context"
	"log"

	models "github.com/imirjar/rb-auth/internal/entity/models"
)

// type iStorage interface{

// }

// Service layer realize authorization, register and authentication methods
type service struct {
	// storage iStorage
}

// return JWT token
func (s *service) Authenticate(ctx context.Context, user models.User) {
	log.Printf("Authenticate: %s", user.ID)
}
func (s *service) Registrate(ctx context.Context, user models.User) {
	log.Printf("Registrate: %s", user.ID)
}
func (s *service) Authorize(ctx context.Context, user models.User) {
	log.Printf("Authorize: %s", user.ID)
}

func New() (*service, error) {
	return &service{}, nil
}
