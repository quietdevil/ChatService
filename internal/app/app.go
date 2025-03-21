package app

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/quietdevil/ChatSevice/internal/config"
	"github.com/quietdevil/ChatSevice/internal/interceptor"
	"github.com/quietdevil/ChatSevice/internal/interceptor/authorization"
	"github.com/quietdevil/ChatSevice/internal/logger"
	"github.com/quietdevil/ChatSevice/pkg/chat_v1"
	closer "github.com/quietdevil/Platform_common/pkg/closer"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	ServiceProvider *ServiceProvider
	ServerGRPC      *grpc.Server
	HttpServer      *http.Server
	AuthGrpcClient  *authorization.ClientAuth
}

func NewApp(cxt context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(cxt)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) initServProv(_ context.Context) error {
	serv := NewServiceProvider()
	a.ServiceProvider = serv
	return nil
}

func (a *App) initHttpServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := chat_v1.RegisterChatHandlerFromEndpoint(ctx, mux, a.ServiceProvider.GRPCConfig().Address(), opts)

	if err != nil {
		return err
	}
	a.HttpServer = &http.Server{
		Addr:    a.ServiceProvider.HttpConfig().Address(),
		Handler: mux,
	}
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := config.Load(); err != nil {
		return err
	}
	return nil
}

func (a *App) initGrpcServer(ctx context.Context) error {
	a.ServerGRPC = grpc.NewServer(grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			a.ServiceProvider.AuthClient(ctx).InterceptorAuthorization,
			interceptor.InterceptorLogger),
		),
	)
	reflection.Register(a.ServerGRPC)
	chat_v1.RegisterChatServer(a.ServerGRPC, a.ServiceProvider.ImplementationChat(ctx))
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	f := []func(ctx context.Context) error{
		a.initConfig,
		a.initServProv,
		a.initGrpcServer,
		a.initHttpServer,
		a.initLogger,
	}

	for _, fu := range f {
		err := fu(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) Run() (err error) {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		defer wg.Done()
		err = a.RunGRPCServer()

	}()

	go func() {
		defer wg.Done()
		err = a.RunHttpServer()

	}()
	log.Print("Server Run")
	wg.Wait()
	return nil
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

func (a *App) RunHttpServer() error {
	if err := a.HttpServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (a *App) initLogger(_ context.Context) error {
	atom := zap.NewAtomicLevel()
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.Lock(os.Stdout), atom)
	core = zapcore.NewSamplerWithOptions(core, time.Second, 3, 3)
	logger.Init(core)
	return nil
}
