package rpc

import (
	"context"
	"github.com/quietdevil/ServiceAuthentication/pkg/access_v1"
)

type ClientRPC struct {
	clientGrpc access_v1.AccessV1Client
}

func NewClientRPC(client access_v1.AccessV1Client) ClientGrpcV1 {
	return &ClientRPC{clientGrpc: client}
}

func (g *ClientRPC) Check(ctx context.Context, endpoint string) error {
	_, err := g.clientGrpc.Check(ctx, &access_v1.CheckRequest{EndpointAddress: endpoint})
	if err != nil {
		return err
	}
	return nil
}
