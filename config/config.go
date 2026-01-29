package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	Port   string `env:"PORT"`
	DBConn string `env:"DB_CONN"`
	Secret string `env:"SECRET"`
}

func New() Config {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Printf("Configuration err %+v\n", err)
	}

	return cfg
}
