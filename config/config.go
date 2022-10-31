package config

import (
	"errors"
	"os"
	"time"
)

type AppConfig struct {
	SqlDb      string
	Retries    int
	RetryAfter time.Duration
}

func NewConfig() (*AppConfig, error) {
	db, ok := os.LookupEnv("SQLITE_DB_NAME")
	if !ok {
		return nil, errors.New("sqlite db name must be set")
	}
	return &AppConfig{SqlDb: db}, nil
}
