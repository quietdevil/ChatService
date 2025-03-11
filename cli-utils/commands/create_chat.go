package Cmds

import (
	"context"
	"fmt"
	"github.com/quietdevil/ChatSevice/cli-utils/root"
	"github.com/quietdevil/ChatSevice/cli-utils/utils"
	"github.com/quietdevil/ChatSevice/pkg/chat_v1"
	"github.com/spf13/cobra"
)

func InitCreateChatCmd(ctx context.Context, client chat_v1.ChatClient) {
	var createChatCmd = &cobra.Command{
		Use:   "create-chat",
		Short: "Create a new chat",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, err := utils.NewContextOutGoing(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
			resp, err := client.Create(ctx, &chat_v1.CreateRequest{})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Successfully created chat [id]:", resp.GetId())
			}

		},
	}

	root.RootCmd.AddCommand(createChatCmd)
}
