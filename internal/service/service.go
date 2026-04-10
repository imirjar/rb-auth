package service

import (
	"context"
	"errors"
	"time"

	"github.com/imirjar/rb-auth/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Storager interface {
	FindByEmail(context.Context, string) (models.User, error)
	AddUser(context.Context, models.User) error
}

// Service layer
// User service: authentificate authorize identificate
// Token service: generate refresh validate
type Service struct {
	Secret  string
	Storage Storager
}

func New(secret string) (*Service, error) {
	return &Service{
		Secret: secret,
	}, nil
}

func (s *Service) CreateUser(ctx context.Context, req domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 4. Создание пользователя для сохранения
	user := models.User{
		Username:     req.Username,
		Email:        req.Email,
		Roles:        []string{},
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	// 5. Сохранение в storage
	err = s.Storage.AddUser(ctx, user)
	if err != nil {
		return err
	}

	return nil

}

func (s *Service) LogIn(ctx context.Context, login domain.LoginRequest) (entities.TokenPair, error) {
	user, err := s.Storage.FindByEmail(ctx, login.Email)
	if err != nil {
		return entities.TokenPair{}, err
	}

	// if user.PasswordHash != string(hashedPassword) {
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(login.Password)); err != nil {
		// log.Println(user.PasswordHash, login.Password)
		return entities.TokenPair{}, errors.New("forbidden")
	}

	// log.Print(user)
	return entities.TokenPair{
		Access:  "xxxx.yyyy.zzzz",
		Refresh: "xxxx.yyyy.zzzz",
	}, nil

}

// return JWT token with prolongated date
func (s *Service) Refresh(ctx context.Context, token string) (entities.TokenPair, error) {
	return entities.TokenPair{
		Access:  "xxxx.yyyy.zzzz",
		Refresh: "xxxx.yyyy.zzzz",
	}, nil
}

func (s *Service) Validate(ctx context.Context, token string) (bool, error) {
	return true, nil
}
