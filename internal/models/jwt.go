package models

import (
	"encoding/json"
	"fmt"
	"time"
)

const TOKEN_EXP = time.Hour * 3
const SECRET_KEY = "supersecretkey"

type JWT struct {
	Header  Header `json:"header"`            // Заголовок
	Payload Claims `json:"payload",omitempty` // Полезная нагрузка
	// Подпись
}

type Header struct {
	Algoritm string `json:"alg"`
	Type     string `json:"typ"`
}

type Claims struct {
	User User  `json:"user,omitempty"`
	Exp  int64 `json:"exp,omitempty"`
}

func (j *JWT) GetSignature() (string, error) {
	// Set header
	j.Header.Algoritm = ""
	j.Header.Algoritm = ""

	expTime := time.Now().Add(time.Hour * 12) // 12 hours
	j.Payload.Exp = expTime.Unix()

	b, err := json.Marshal(j)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "", err
	}

	// h := sha256.New()
	// h.Write([]byte(b))
	// // возвращаем строку токена

	return string(b), nil
}
