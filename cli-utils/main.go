package main

import (
	"github.com/quietdevil/ChatSevice/cli-utils/cmd"
	"github.com/quietdevil/ChatSevice/cli-utils/cmd/commands"
)

func main() {
	rootCmd := cmd.RootCmd()
	commands.InitLogin(rootCmd)
	rootCmd.Execute()

}
