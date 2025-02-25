package config

import (
	"errors"
	"os"
)

type OpensslConfig interface {
	PathPem() string
	PathKey() string
}

type opensslConfig struct {
	pathPem string
	pathKey string
}

func NewOpensslConfig() (OpensslConfig, error) {
	pathPem := os.Getenv("PATH_PEM")
	if pathPem == "" {
		return nil, errors.New("don't parse path service.pem")
	}
	pathKey := os.Getenv("PATH_KEY")
	if len(pathKey) == 0 {
		return nil, errors.New("don't parse path service.key")
	}
	return &opensslConfig{pathKey: pathKey, pathPem: pathPem}, nil
}

func (o *opensslConfig) PathPem() string {
	return o.pathPem
}

func (o *opensslConfig) PathKey() string {
	return o.pathKey
}
