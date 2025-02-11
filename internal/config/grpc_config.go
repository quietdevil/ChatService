package config

import (
	"errors"
	"net"
	"os"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGrpcConfig() (GRPCConfig, error) {
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		return nil, errors.New("don't parse grpc port")
	}
	host := os.Getenv("GRPC_HOST")
	if host == "" {
		return nil, errors.New("don't parse grpc host")
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (g *grpcConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}
