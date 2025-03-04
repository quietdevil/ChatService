package main

import (
	"chatservice/cli-utils/cmd"
	"chatservice/cli-utils/cmd/commands"
)

func main() {
	rootCmd := cmd.RootCmd()
	commands.InitLogin(rootCmd)
	rootCmd.Execute()

}
