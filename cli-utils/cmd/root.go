package cmd

import (
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	var (
		rootCmd = &cobra.Command{
			Use:   "tiena",
			Short: "Tiena is a CLI tool to manage chats",
			Long:  `A small program for communicating via a web server, where you and your friends can discuss dirty business`,
			Run: func(cmd *cobra.Command, args []string) {
				// Do Stuff Here
			},
		}
	)
	return rootCmd
}
