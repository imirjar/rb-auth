package service

import (
	"context"
	"log"

	models "github.com/imirjar/rb-auth/internal/entity/models"
)

type storage interface {
	AddUser(models.User) error
	GetUser(string) (models.User, error)
}

// Service layer realize authorization, register and authentication methods
type service struct {
	Storage storage
}

// return JWT token
func (s *service) Authenticate(ctx context.Context, user models.User) (models.User, error) {
	return s.Storage.GetUser(user.Login)
}
func (s *service) Registrate(ctx context.Context, user models.User) error {
	err := s.Storage.AddUser(user)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("Registrate: %s", user.ID)
	return nil
}
func (s *service) Authorize(ctx context.Context, user models.User) {
	log.Printf("Authorize: %s", user.ID)
}

func New() (*service, error) {
	return &service{}, nil
}
