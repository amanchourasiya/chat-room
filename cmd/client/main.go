package main

import (
	"fmt"

	"github.com/amanchourasiya/chat-room/cmd/client/cmd"
	"github.com/amanchourasiya/chat-room/internal/client"
	"github.com/spf13/cobra"
)

func serve(_ *cobra.Command, _ []string) error {
	clnt := &client.Client{}
	if cmd.Port != 0 {
		fmt.Printf("Entered port %d\n", cmd.Port)
		clnt.Port = cmd.Port
	}
	if cmd.ServerName != "" {
		fmt.Printf("Entered server name %s\n", cmd.ServerName)
		clnt.Name = cmd.ServerName
	}
	if cmd.ServerAddr != "" {
		fmt.Printf("Connecting to server: %s\n", cmd.ServerAddr)
		clnt.ServerAddr = cmd.ServerAddr
	}

	clnt.Start()

	return nil
}

func main() {
	cmd.RootCmd.RunE = serve
	cmd.RootCmd.Execute()
}
