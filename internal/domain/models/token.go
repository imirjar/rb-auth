package models

import (
	"github.com/golang-jwt/jwt/v5"
)

// Claims - кастомные claims для нашего приложения
type Claims struct {
	jwt.RegisteredClaims

	// User identity
	UserID   string   `json:"user_id"`
	Email    string   `json:"email"`
	Username string   `json:"username,omitempty"`
	Roles    []string `json:"roles,omitempty"`

	// Additional context
	SessionID string                 `json:"session_id,omitempty"`
	DeviceID  string                 `json:"device_id,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

// Конвертер Claims → UserProfile (если нужно)
func (c *Claims) ToProfile() *UserProfile {
	return &UserProfile{
		ID:       c.UserID,
		Email:    c.Email,
		Username: c.Username,
		Roles:    c.Roles,
	}
}
