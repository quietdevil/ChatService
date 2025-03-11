package main

import (
	"context"
	"github.com/quietdevil/ChatSevice/cli-utils/commands"
	"github.com/quietdevil/ChatSevice/cli-utils/root"
	"github.com/quietdevil/ChatSevice/pkg/chat_v1"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_user_v1"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

func main() {
	ctx := context.Background()
	tramCredis, err := credentials.NewClientTLSFromFile("/home/ivan/Documents/GolandProjects/ChatService/keys/service.pem", "")
	if err != nil {

		log.Fatal(err)
	}

	conn, err := grpc.NewClient(net.JoinHostPort("localhost", "40000"), grpc.WithTransportCredentials(tramCredis))
	if err != nil {
		log.Fatal(err)
	}
	clientAuth := auth_v1.NewAuthenticationV1Client(conn)

	connChat, err := grpc.NewClient(net.JoinHostPort("localhost", "50000"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	clientChat := chat_v1.NewChatClient(connChat)

	connUser, err := grpc.NewClient(net.JoinHostPort("localhost", "40000"), grpc.WithTransportCredentials(tramCredis))
	if err != nil {
		log.Fatal(err)
	}
	clientUser := auth_user_v1.NewAuthenticationUserV1Client(connUser)

	Cmds.InitLoginCmd(ctx, clientAuth)
	Cmds.InitCreateChatCmd(ctx, clientChat)
	Cmds.InitRegistrationCmd(ctx, clientUser)
	Cmds.InitConnectionChatCmd(ctx, clientChat)

	root.RootCmd.Execute()

}
