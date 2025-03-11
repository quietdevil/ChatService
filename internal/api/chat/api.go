package chat

import (
	"context"
	"github.com/quietdevil/ChatSevice/internal/service"
	desc "github.com/quietdevil/ChatSevice/pkg/chat_v1"
	"sync"
)

type ImplementationChat struct {
	desc.UnimplementedChatServer

	chats map[string]*Chat
	mx    sync.RWMutex

	channels    map[string]chan *desc.Message
	mxChannels  sync.RWMutex
	ChatService service.Service
}

func NewImplementation(ctx context.Context, service service.Service) *ImplementationChat {
	return &ImplementationChat{
		chats:       make(map[string]*Chat),
		channels:    make(map[string]chan *desc.Message),
		ChatService: service,
	}
}

type Chat struct {
	streams map[string]desc.Chat_ConnectServer
	m       sync.RWMutex
}
