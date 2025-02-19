package authorization

import (
	"chatservice/internal/client/rpc"
	"context"
	"errors"
	"fmt"
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
		return nil, errors.New("the token did not arrive")
	}
	token, ok := md["authorization"]
	fmt.Println(token, info.FullMethod)
	if !ok {
		return nil, errors.New("the token did not arrive")
	}
	newMd := metadata.New(map[string]string{"authorization": token[0]})
	ctx = metadata.NewOutgoingContext(ctx, newMd)
	err = c.clientAuth.Check(ctx, info.FullMethod)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)

}
