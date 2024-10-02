package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type (
	// App configuration from config.yml
	Config struct {
		Http     Http     `yaml:"http"`
		Security Security `yaml:"crypto"`
	}
	// Configuration for http server only
	Http struct {
		Port string `yaml:"port"`
	}
	Security struct {
		Pub  Pub  `yaml:"public_key"`
		Priv Priv `yaml:"private_key"`
	}
	Priv struct {
		path string
		Key  *rsa.PrivateKey
	}
	Pub struct {
		path string
		Key  *rsa.PublicKey
	}
)

func (s *Pub) UnmarshalYAML(keyPath *yaml.Node) error {

	if err := keyPath.Decode(&s.path); err != nil {
		return err
	}

	if s.path != "" {
		// log.Println("###UNMARSH PUB", s.path)
		keyData, err := os.ReadFile(s.path)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(keyData)
		if block == nil || block.Type != "RSA PUBLIC KEY" {
			log.Println("pub block err:", err)
			log.Println("pub type:", block.Type)
			return err
		}

		pk, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			log.Println("###", err)
			return err
		}

		s.Key = pk
	}

	return nil
}

func (s *Priv) UnmarshalYAML(value *yaml.Node) error {

	if err := value.Decode(&s.path); err != nil {
		return err
	}

	if s.path != "" {
		// log.Println("###UNMARSH PRIV", s.path)
		keyData, err := os.ReadFile(s.path)
		if err != nil {
			return err
		}

		block, _ := pem.Decode(keyData)
		if block == nil || block.Type != "RSA PRIVATE KEY" {
			return err
		}

		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return err
		}

		s.Key = privateKey
	}

	return nil
}

func New() *Config {
	var cfg Config

	file, err := os.ReadFile("config/config.yml")
	if err != nil {
		log.Fatalf("error opening config file: %v", err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// log.Print(cfg.Security.Pub)
	// log.Print(cfg.Security.Priv)
	return &cfg
}
