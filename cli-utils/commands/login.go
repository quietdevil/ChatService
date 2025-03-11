package Cmds

import (
	"context"
	"fmt"
	"github.com/quietdevil/ChatSevice/cli-utils/models"
	"github.com/quietdevil/ChatSevice/cli-utils/root"
	"github.com/quietdevil/ChatSevice/cli-utils/utils"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_v1"
	"github.com/spf13/cobra"
)

func InitLoginCmd(ctx context.Context, client auth_v1.AuthenticationV1Client) {
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to Tiena",
		Run: func(cmd *cobra.Command, args []string) {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			respRefresh, _ := client.Login(ctx, &auth_v1.LoginRequest{Username: username, Password: password})
			respAccessToken, err := client.GetAccessToken(ctx, &auth_v1.GetAccessTokenRequest{RefreshToken: respRefresh.GetRefreshToken()})
			if err != nil {
				fmt.Println("No logged because -> ", err)
				return
			} else {
				login := &models.Login{
					AccessToken:  respAccessToken.GetAccessToken(),
					RefreshToken: respRefresh.GetRefreshToken(),
				}

				err = utils.MarshalTokensInFile(login, "login.json")
				if err != nil {
					fmt.Println("No logged because -> ", err)
					return
				}
				fmt.Println("Successfully logged in")
			}

		},
	}
	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
	root.RootCmd.AddCommand(loginCmd)

}
