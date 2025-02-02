package main

import (
	desc "chatservice/pkg/chat_v1"
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = "40000"

type server struct {
	desc.UnimplementedChatServer
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponce, error) {
	return &desc.CreateResponce{
		Id: 64,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	message := fmt.Sprintf("[from] %v [message]: %v", req.From, req.Text)
	fmt.Fprintln(os.Stdout, message)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+grpcPort)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatServer(s, &server{})

	log.Printf("server listing %v, port %v", lis.Addr().Network(), grpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
