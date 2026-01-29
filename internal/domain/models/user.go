package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           string     `json:"id" db:"id"`
	Email        string     `json:"email" db:"email"`
	Username     string     `json:"username" db:"username"`
	PasswordHash string     `json:"-" db:"password_hash"`
	IsActive     bool       `json:"is_active" db:"is_active"`
	Roles        []string   `json:"roles" db:"roles"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
}

// Конвертер User → UserProfile (для API ответов)
func (u *User) ToProfile() *UserProfile {
	return &UserProfile{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
		Roles:    u.Roles,
	}
}

// Конвертер User → Claims (для JWT)
func (u *User) ToClaims(expiresIn time.Duration) *Claims {
	return &Claims{
		UserID:   u.ID,
		Email:    u.Email,
		Username: u.Username,
		Roles:    u.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   u.ID,
			Issuer:    "your-app",
			Audience:  []string{"api"},
		},
	}
}
