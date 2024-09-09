package storage

import "github.com/imirjar/rb-auth/internal/models"

func New() (*storage, error) {
	return &storage{}, nil
}

type storage struct {
	Users  []models.User
	Groups []models.Group
	Roles  []models.Role
}
