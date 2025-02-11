package config

import (
	"errors"
	"os"
)

type PGConfig interface {
	DSN() string
}

type DBConfig struct {
	Dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		return nil, errors.New("don't parse dsn")
	}
	return &DBConfig{Dsn: dsn}, nil
}

func (c *DBConfig) DSN() string {
	return c.Dsn
}
