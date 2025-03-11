package Cmds

import (
	"context"
	"fmt"
	"github.com/quietdevil/ChatSevice/cli-utils/root"
	"github.com/quietdevil/ServiceAuthentication/pkg/auth_user_v1"
	"github.com/spf13/cobra"
)

func InitRegistrationCmd(ctx context.Context,
	clientUser auth_user_v1.AuthenticationUserV1Client,
) {
	var regisCmd = &cobra.Command{
		Use:   "registration",
		Short: "Register a user",
		Long:  "A username and password are required for registration.",
		Run: func(cmd *cobra.Command, args []string) {
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			email, _ := cmd.Flags().GetString("email")
			passwordConfirmation, _ := cmd.Flags().GetString("password_confirm")
			role, _ := cmd.Flags().GetString("role")

			roleNum, ok := auth_user_v1.Enum_value[role]
			if !ok {
				fmt.Println("Role does not exist")
				return
			}

			_, err := clientUser.Create(ctx, &auth_user_v1.CreateRequest{
				Name:            username,
				Email:           email,
				Password:        password,
				PasswordConfirm: passwordConfirmation,
				Role:            auth_user_v1.Enum(roleNum),
			},
			)
			if err != nil {
				fmt.Println("Registration Failed: ", err)
				return
			}

			fmt.Println("Successfully registered user", username)

		},
	}

	root.RootCmd.AddCommand(regisCmd)
	regisCmd.Flags().StringP("username", "u", "", "Username")
	regisCmd.Flags().StringP("email", "e", "", "Email")
	regisCmd.Flags().StringP("password", "p", "", "Password")
	regisCmd.Flags().StringP("password_confirm", "c", "", "PasswordConfirm")
	regisCmd.Flags().StringP("role", "r", "", "Role")
	regisCmd.MarkFlagsMutuallyExclusive("role")

}
