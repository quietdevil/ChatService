package config

import (
	"log"
	"net"
	"os"
)

type HTTPConfig interface {
	Address() string
}

type HttpConfig struct {
	host string
	port string
}

func NewHttpConfig() HTTPConfig {
	host := os.Getenv("HTTP_HOST")
	if len(host) == 0 {
		log.Fatal("don't parse http host")
	}

	port := os.Getenv("HTTP_PORT")
	if len(port) == 0 {
		log.Fatal("don't parse http port")
	}

	return &HttpConfig{
		host: host,
		port: port,
	}
}

func (h *HttpConfig) Address() string {
	return net.JoinHostPort(h.host, h.port)
}
