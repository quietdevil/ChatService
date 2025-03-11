package Cmds

import (
	"bufio"
	"context"
	"fmt"
	"github.com/quietdevil/ChatSevice/cli-utils/root"
	"github.com/quietdevil/ChatSevice/cli-utils/utils"
	"github.com/quietdevil/ChatSevice/pkg/chat_v1"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"io"
	"os"
)

func InitConnectionChatCmd(ctx context.Context, client chat_v1.ChatClient) {
	var connectCmd = &cobra.Command{
		Use:   "connect",
		Short: "Connect to a chat by id",
		Long:  `Connect to a chat by id`,
		Run: func(cmd *cobra.Command, args []string) {

			id, _ := cmd.Flags().GetString("id")

			loginSettings, err := utils.UnmarshalTokensInFile()
			if err != nil {
				fmt.Println(err)
				return
			}
			username, err := utils.UsernameFromAccessToken(loginSettings.AccessToken, "tennis")
			if err != nil {
				fmt.Println(err)
				return
			}
			ctxNew, err := utils.NewContextOutGoing(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
			stream, err := client.Connect(ctxNew, &chat_v1.ConnectRequest{Id: id, Username: username})
			fmt.Println(err)
			if err != nil {
				fmt.Println("Error connecting:", err)
				fmt.Println(err)
				return
			}
			fmt.Println("Successfully connected to the chat")
			go receiver(stream)
			for {
				sc := bufio.NewScanner(os.Stdin)
				fmt.Print("> ")
				for sc.Scan() {
					txt := sc.Text()
					if txt == "" {
						fmt.Println("Message cannot be empty!")
					}
					_, err = client.SendMessage(ctxNew, &chat_v1.SendMessageRequest{Id: id, Message: &chat_v1.Message{
						From: username,
						Text: txt}})
					if err != nil {
						fmt.Println(err)
						return
					}
				}
			}
		},
	}
	connectCmd.Flags().StringP("id", "i", "", "Chat ID")
	root.RootCmd.AddCommand(connectCmd)
}

func receiver(stream grpc.ServerStreamingClient[chat_v1.Message]) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed.")
				return
			}
			fmt.Println(err)
			return
		}
		fmt.Printf("From: %v Time: %v - Received message: %v\n", msg.GetFrom(), msg.GetTimestamp(), msg.GetText())
		fmt.Println("> ")
	}

}
