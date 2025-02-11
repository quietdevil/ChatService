package chat

import (
	desc "chatservice/pkg/chat_v1"
	"context"
	"fmt"
	"os"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponce, error) {
	id, err := s.ChatService.Create(ctx, req.Usernames)
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponce{
		Id: int64(id),
	}, nil
}

func (s *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.ChatService.Delete(ctx, int(req.GetId()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	message := fmt.Sprintf("[from] %v [message]: %v", req.From, req.Text)
	fmt.Fprintln(os.Stdout, message)
	return &emptypb.Empty{}, nil
}
