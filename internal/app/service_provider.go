package app

import (
	api "chatservice/internal/api/chat"
	"chatservice/internal/client/rpc"
	"chatservice/internal/config"
	"chatservice/internal/interceptor/authorization"
	"chatservice/internal/repository"
	"chatservice/internal/repository/chat"
	"chatservice/internal/repository/logs"
	"chatservice/internal/service"
	"chatservice/internal/service/chats"
	"context"
	"github.com/quietdevil/ServiceAuthentication/pkg/access_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"

	closer "github.com/quietdevil/Platform_common/pkg/closer"
	db "github.com/quietdevil/Platform_common/pkg/db"
	pg "github.com/quietdevil/Platform_common/pkg/db/pg"
	transaction "github.com/quietdevil/Platform_common/pkg/db/transaction"
)

type ServiceProvider struct {
	PgConfig         config.PGConfig
	GrpcConfig       config.GRPCConfig
	HTTPConfig       config.HTTPConfig
	GRPCClientConfig config.GRPCClientConfig
	opensslConfig    config.OpensslConfig
	ClientAuth       *authorization.ClientAuth
	ClientDB         db.Client
	txManager        db.TxManager
	clientGrpc       rpc.ClientGrpcV1
	logger           repository.Logger
	Repository       repository.Repository
	Service          service.Service
	Implementation   *api.Implementation
}

func NewServiceProvider() *ServiceProvider {

	return &ServiceProvider{}
}

func (s *ServiceProvider) GrpcClientConfig() config.GRPCClientConfig {
	if s.GRPCClientConfig == nil {
		c, err := config.NewGrpcClientConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.GRPCClientConfig = c
	}
	return s.GRPCClientConfig

}

func (s *ServiceProvider) OpenSSLConfig() config.OpensslConfig {
	if s.opensslConfig == nil {
		c, err := config.NewOpensslConfig()
		if err != nil {
			log.Fatal(err)
		}
		s.opensslConfig = c
	}
	return s.opensslConfig
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

func (s *ServiceProvider) HttpConfig() config.HTTPConfig {
	if s.HTTPConfig == nil {
		s.HTTPConfig = config.NewHttpConfig()
	}
	return s.HTTPConfig
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

func (s *ServiceProvider) Logger(ctx context.Context) repository.Logger {
	if s.logger == nil {
		l := logs.NewLogger(s.ClientDb(ctx))
		s.logger = l
	}
	return s.logger
}

func (s *ServiceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		txMan := transaction.NewManager(s.ClientDb(ctx).DB())
		s.txManager = txMan
	}
	return s.txManager
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
		serv := chats.NewService(s.ChatRepository(ctx), s.TxManager(ctx), s.Logger(ctx))
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

func (s *ServiceProvider) ClientGrpc(ctx context.Context) rpc.ClientGrpcV1 {
	if s.clientGrpc == nil {
		transportCreds, err := credentials.NewClientTLSFromFile(s.OpenSSLConfig().PathPem(), "")
		if err != nil {
			log.Fatal(err)
		}
		conn, err := grpc.NewClient(s.GrpcClientConfig().Address(), grpc.WithTransportCredentials(transportCreds))
		if err != nil {
			log.Fatal(err)
		}
		accessClient := access_v1.NewAccessV1Client(conn)
		s.clientGrpc = rpc.NewClientRPC(accessClient)
	}
	return s.clientGrpc
}

func (s *ServiceProvider) AuthClient(ctx context.Context) *authorization.ClientAuth {
	if s.ClientAuth == nil {
		c := authorization.NewClientAuth(s.ClientGrpc(ctx))
		s.ClientAuth = c
	}
	return s.ClientAuth
}
