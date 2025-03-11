package authorization

import (
	"context"
	"errors"
	"github.com/quietdevil/ChatSevice/internal/client/rpc"
	"github.com/quietdevil/ChatSevice/internal/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type ClientAuth struct {
	clientAuth rpc.ClientGrpcV1
}

func NewClientAuth(cln rpc.ClientGrpcV1) *ClientAuth {
	return &ClientAuth{clientAuth: cln}
}

func (c *ClientAuth) InterceptorAuthorization(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		logger.Error("the token did not arrive")
		return nil, errors.New("the token did not arrive")
	}
	token, ok := md["authorization"]
	if !ok {
		logger.Error("the token did not arrive")
		return nil, errors.New("the token did not arrive")
	}
	logger.Info("the token is " + token[0])
	newMd := metadata.New(map[string]string{"authorization": token[0]})
	ctx = metadata.NewOutgoingContext(ctx, newMd)
	err = c.clientAuth.Check(ctx, info.FullMethod)
	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}
	return handler(ctx, req)

}
