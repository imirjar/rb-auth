package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "supersecretkey"

type Claims struct {
	User User `json:"user,omitempty"`
	jwt.RegisteredClaims
}

type JWT struct {
	Access  Claims `json:"access,omitempty"`
	Refresh Claims `json:"refresh,omitempty"`
}
