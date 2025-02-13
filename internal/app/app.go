package app

import (
	"chatservice/internal/config"
	"chatservice/pkg/chat_v1"
	"context"
	"log"
	"net"

	closer "github.com/quietdevil/Platform_common/pkg/closer"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	ServiceProvider *ServiceProvider
	ServerGRPC      *grpc.Server
}

func NewApp(cxt context.Context) *App {
	app := &App{}
	app.initDeps(cxt)

	return app
}

func (a *App) initServProv(ctx context.Context) error {
	serv := NewServiceProvider()
	a.ServiceProvider = serv
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := config.Load(); err != nil {
		return err
	}
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.ServerGRPC = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	reflection.Register(a.ServerGRPC)
	chat_v1.RegisterChatServer(a.ServerGRPC, a.ServiceProvider.ImplementationChat(ctx))
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	f := []func(ctx context.Context) error{
		a.initConfig,
		a.initServProv,
		a.initGrpcServer,
	}

	for _, fu := range f {
		fu(ctx)
	}
	return nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	log.Print("Server Run")
	return a.RunGRPCServer()
}

func (a *App) RunGRPCServer() error {
	lis, err := net.Listen("tcp", a.ServiceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	if err := a.ServerGRPC.Serve(lis); err != nil {
		return err
	}
	return nil
}
