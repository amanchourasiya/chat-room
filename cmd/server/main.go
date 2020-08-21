package main

import (
	"fmt"

	"github.com/amanchourasiya/chat-room/cmd/client/cmd"
	"github.com/spf13/cobra"
)

func serve(_ *cobra.Command, _ []string) error {
	if cmd.Port != 0 {
		fmt.Printf("Entered port %d\n", cmd.Port)
	}
	if cmd.ServerName != "" {
		fmt.Printf("Entered server name %s\n", cmd.ServerName)
	}
	return nil
}

func main() {
	cmd.RootCmd.RunE = serve
	cmd.RootCmd.Execute()
}
