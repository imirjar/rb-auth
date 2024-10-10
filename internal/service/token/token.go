package token

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/imirjar/rb-auth/internal/models"
)

// Service layer
type service struct {
	PrivKey *rsa.PrivateKey
	PubKey  *rsa.PublicKey
}

// return JWT token
func (s *service) Create(ctx context.Context, user models.User) (string, error) {
	claims := models.Claims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(8 * time.Hour)),
		},
	}
	// Watch out! Signiing method for RS256
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(s.PrivKey)
	if err != nil {
		// log.Println(err, reflect.TypeOf(s.SignKey), s.SignKey)
		return "", err
	}

	return ss, nil
}

// return JWT token with prolongated date
func (s *service) Refresh(ctx context.Context, token string) (string, error) {
	// read old get claims
	// change experied
	// return new
	return token, nil
}

func (s *service) Read(ctx context.Context, tokenString string) (models.User, error) {

	claims := models.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.PubKey, nil
	})
	if err != nil {
		log.Print(err)
		return claims.User, err
	}

	if !token.Valid {
		return claims.User, fmt.Errorf("token is not Valid")
	}

	return claims.User, nil
}

func (s *service) Validate(ctx context.Context, tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.PubKey, nil
	})
	if err != nil {
		log.Print("Parse error")
		log.Print(err)
		return false
	}

	if !token.Valid {
		log.Print("invalid ebany")
		return false
	}

	return true
}

func New(priv *rsa.PrivateKey, pub *rsa.PublicKey) (*service, error) {
	// log.Print(priv)
	// log.Print(pub)
	return &service{
		PubKey:  pub,
		PrivKey: priv,
	}, nil
}
