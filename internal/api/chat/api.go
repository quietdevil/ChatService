package chat

import (
	"chatservice/internal/service"
	desc "chatservice/pkg/chat_v1"
	"context"
)

type Implementation struct {
	desc.UnimplementedChatServer
	ChatService service.Service
}

func NewImplementation(ctx context.Context, service service.Service) *Implementation {
	return &Implementation{
		ChatService: service,
	}
}
