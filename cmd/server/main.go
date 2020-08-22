package main

import (
	"fmt"

	"github.com/amanchourasiya/chat-room/cmd/server/cmd"
	"github.com/amanchourasiya/chat-room/internal/server"
	"github.com/spf13/cobra"
)

func serve(_ *cobra.Command, _ []string) error {
	srv := &server.Server{}
	if cmd.Port != 0 {
		fmt.Printf("Entered port %d\n", cmd.Port)
		srv.Port = cmd.Port
	}
	if cmd.ServerName != "" {
		fmt.Printf("Entered server name %s\n", cmd.ServerName)
		srv.Name = cmd.ServerName
	}
	srv.Start()

	return nil
}

func main() {
	cmd.RootCmd.RunE = serve
	cmd.RootCmd.Execute()
}
