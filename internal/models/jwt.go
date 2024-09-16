package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "supersecretkey"

type JWT struct {
	Header  Header `json:"header"`  // Заголовок
	Payload Claims `json:"payload"` // Полезная нагрузка
	// Подпись
}

type Header struct {
	Algoritm string `json:"alg"`
	Type     string `json:"typ"`
}

type Claims struct {
	jwt.RegisteredClaims
	UserID int
}

func (j *JWT) GetSignature() (string, error) {
	// создаём новый токен с алгоритмом подписи HS256 и утверждениями — Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			// когда создан токен
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXP)),
		},
		// собственное утверждение
		UserID: 1,
	})

	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	// возвращаем строку токена
	return tokenString, nil
}
