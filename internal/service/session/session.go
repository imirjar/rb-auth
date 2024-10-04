package session

import (
	"context"
)

// Service layer
type service struct {
}

func CreateSession(ctx context.Context, token string) error {
	return nil
}

func DeleteSession(ctx context.Context, sessionID int) error {
	return nil
}

func New() (*service, error) {
	return &service{}, nil
}
