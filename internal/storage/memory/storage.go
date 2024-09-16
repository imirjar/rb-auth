package storage

import (
	"log"

	"github.com/imirjar/rb-auth/internal/models"
)

func New() (*storage, error) {
	return &storage{
		Users:  make(map[string]models.User),
		Groups: make(map[string]models.Group),
		Roles:  make(map[string]models.Role),
	}, nil
}

type storage struct {
	Users  map[string]models.User
	Groups map[string]models.Group
	Roles  map[string]models.Role
}

func (s *storage) AddUser(user models.User) error {
	if _, ok := s.Users[user.Login]; !ok {
		s.Users[user.Login] = user
		return nil
	} else {
		return errUserExists
	}
}

func (s *storage) GetUser(login string) (models.User, error) {
	log.Println("STORAGE", s.Users)

	if user, ok := s.Users[login]; !ok {
		return user, nil
	} else {
		return user, errFindUser
	}

}
