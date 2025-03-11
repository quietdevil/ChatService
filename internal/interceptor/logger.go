package interceptor

import (
	"context"
	"github.com/quietdevil/ChatSevice/internal/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

func InterceptorLogger(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	timeNow := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error(), zap.String("method", info.FullMethod), zap.Any("req", res))
	}
	logger.Info("request", zap.String("method", info.FullMethod), zap.Any("req", res), zap.Duration("duration", time.Since(timeNow)))
	return res, nil
}
