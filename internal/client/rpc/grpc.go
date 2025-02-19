package rpc

import "context"

type ClientGrpcV1 interface {
	Check(context.Context, string) error
}
