package app

import (
	api "chatservice/internal/api/chat"
	"chatservice/internal/client/db"
	"chatservice/internal/client/db/pg"
	"chatservice/internal/closer"
	"chatservice/internal/config"
	"chatservice/internal/repository"
	"chatservice/internal/repository/chat"
	"chatservice/internal/service"
	"chatservice/internal/service/chats"
	"context"
	"log"
)

type ServiceProvider struct {
	PgConfig       config.PGConfig
	GrpcConfig     config.GRPCConfig
	ClientDB       db.Client
	Repository     repository.Repository
	Service        service.Service
	Implementation *api.Implementation
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}

func (s *ServiceProvider) PGConfig() config.PGConfig {
	if s.PgConfig == nil {
		c, err := config.NewPGConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.PgConfig = c
	}
	return s.PgConfig
}

func (s *ServiceProvider) GRPCConfig() config.GRPCConfig {
	if s.GrpcConfig == nil {
		c, err := config.NewGrpcConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.GrpcConfig = c
	}
	return s.GrpcConfig
}

func (s *ServiceProvider) ClientDb(ctx context.Context) db.Client {
	if s.ClientDB == nil {
		client, err := pg.NewDBClient(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatal(err)
		}
		err = client.DB().Ping(ctx)
		if err != nil {
			log.Fatal(err)
		}
		closer.Add(client.Close)
		s.ClientDB = client

	}
	return s.ClientDB

}

func (s *ServiceProvider) ChatRepository(ctx context.Context) repository.Repository {
	if s.Repository == nil {
		repos := chat.NewRepository(s.ClientDb(ctx))
		s.Repository = repos
	}
	return s.Repository

}

func (s *ServiceProvider) ChatService(ctx context.Context) service.Service {
	if s.Service == nil {
		serv := chats.NewService(s.ChatRepository(ctx))
		s.Service = serv
	}
	return s.Service
}

func (s *ServiceProvider) ImplementationChat(ctx context.Context) *api.Implementation {
	if s.Implementation == nil {
		server := api.NewImplementation(ctx, s.ChatService(ctx))
		s.Implementation = server
	}
	return s.Implementation
}
