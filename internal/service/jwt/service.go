package service

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

// return JWT token
func (s *service) BuildJWTString(ctx context.Context, user models.User) (string, error) {
	if !user.IsValid() {
		return "", errors.New("user isn't valid")
	}

	user, err := s.Storage.GetUser(user.Login)
	if err != nil {
		log.Println("###", err)
		return "", err
	}

	jwt := models.JWT{
		Payload: models.Claims{
			User: user,
		},
	}

	jwtstring, err := jwt.GetSignature()
	if err != nil {
		log.Print(err)
		return "", err
	}

	return jwtstring, nil
}

func (s *service) Registrate(ctx context.Context, user models.User) error {
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
