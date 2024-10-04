package user

import (
	"context"
	"errors"
	"log"

	"github.com/imirjar/rb-auth/internal/models"
)

type storage interface {
	AddUser(models.User) error
	GetUser(string) (models.User, error)
}

// Service layer realize authorization, register and authentication methods
type service struct {
	Storage storage
}

// return (user, true) is user is exists
func (s *service) CheckUser(ctx context.Context, user models.User) (models.User, bool) {
	if !user.IsValid() {
		log.Println("CheckUser ERROR: user ISN'T VALID")
		return user, false
		// errors.New("user isn't valid")
	}

	user, err := s.Storage.GetUser(user.Login)
	if err != nil {
		log.Println("GetUser ERROR:", err)
		return user, false
	}

	return user, true
}

func (s *service) AddUser(ctx context.Context, user models.User) error {
	if !user.IsValid() {
		return errors.New("user isn't valid")
	}

	err := s.Storage.AddUser(user)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Printf("Registrate: %v", user.ID)
	return nil
}

func New() (*service, error) {
	return &service{}, nil
}
