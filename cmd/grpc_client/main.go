package main

import (
	"context"
	"github.com/quietdevil/ChatSevice/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	conn, err := grpc.NewClient(":50000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := chat_v1.NewChatClient(conn)
	resp, err := client.Create(ctx, &chat_v1.CreateRequest{})
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Connect(ctx, &chat_v1.ConnectRequest{Id: resp.String(), Username: "Ivan"})
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(time.Second * 3)

	msg := chat_v1.Message{
		From: "Ivan",
		Text: "Hello Ivan",
	}

	for {
		select {
		case <-ticker.C:
			_, err := client.SendMessage(ctx, &chat_v1.SendMessageRequest{Message: &msg, Id: resp.String()})
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}
