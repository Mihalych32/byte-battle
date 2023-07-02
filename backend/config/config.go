package config

import (
	"fmt"
	"os"
)

type (
	Config struct {
		DB DB
	}
	DB struct {
		Host         string
		Port         string
		User         string
		Pass         string
		Name         string
		SSLMode      string
		MaxOpenConns int
		MaxIdleConns int
	}
)

func NewConfig() (cfg *Config, err error) {
	cfg = &Config{}

	host := os.Getenv("POSTGRES_HOST")
	if host == "" {
		err = fmt.Errorf("POSTGRES_HOST variable not found in .env file")
		return
	}
	port := os.Getenv("POSTGRES_PORT")
	if port == "" {
		err = fmt.Errorf("POSTGRES_PORT variable not found in .env file")
		return
	}
	user := os.Getenv("POSTGRES_USER")
	if user == "" {
		err = fmt.Errorf("POSTGRES_USER variable not found in .env file")
		return
	}
	name := os.Getenv("POSTGRES_NAME")
	if name == "" {
		err = fmt.Errorf("POSTGRES_NAME variable not found in .env file")
		return
	}
	pass := os.Getenv("POSTGRES_PASSWORD")
	if pass == "" {
		err = fmt.Errorf("POSTGRES_PASSWORD variable not found in .env file")
		return
	}
	sslmode := os.Getenv("POSTGRES_SSLMODE")
	if sslmode == "" {
		err = fmt.Errorf("POSTGRES_SSLMODE variable not found in .env file")
		return
	}

	cfg.DB.Host = host
	cfg.DB.Port = port
	cfg.DB.User = user
	cfg.DB.Name = name
	cfg.DB.Pass = pass
	cfg.DB.SSLMode = sslmode
	return
}
