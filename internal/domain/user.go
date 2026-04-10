package domain

import (
	"time"
)

type User struct {
	ID           string    `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	Login        string    `db:"login" json:"login"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Updated      time.Time `db:"updated_at" json:"updated_at"`
	IsActive     bool      `db:"is_active" json:"is_active"`
}

type Credentials struct {
	ID           string    `db:"id" json:"id"`
	Email        string    `db:"email" json:"email"`
	Login        string    `db:"login" json:"login"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Updated      time.Time `db:"updated_at" json:"updated_at"`
	IsActive     bool      `db:"is_active" json:"is_active"`
}
