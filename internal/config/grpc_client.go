package config

import (
	"errors"
	"net"
	"os"
)

type GRPCClientConfig interface {
	Address() string
}

type grpcClientConfig struct {
	host string
	port string
}

func NewGrpcClientConfig() (GRPCClientConfig, error) {
	port := os.Getenv("GRPCClient_PORT")
	if port == "" {
		return nil, errors.New("don't parse grpc port")
	}
	host := os.Getenv("GRPCClient_HOST")
	if host == "" {
		return nil, errors.New("don't parse grpc host")
	}
	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (g *grpcClientConfig) Address() string {
	return net.JoinHostPort(g.host, g.port)
}
