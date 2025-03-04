package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func InitLogin(root *cobra.Command) {
	var loginCmd = &cobra.Command{
		Use:   "login",
		Short: "Login to Tiena",
		Run: func(cmd *cobra.Command, args []string) {
			username, _ := cmd.Flags().GetString("username")
			fmt.Println("Login called with", username)
		},
	}
	root.AddCommand(loginCmd)
	loginCmd.Flags().StringP("gitusername", "u", "", "Username")
	loginCmd.Flags().StringP("password", "p", "", "Password")

}
