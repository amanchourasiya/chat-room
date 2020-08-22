package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Port       int
	ServerName string
)

var RootCmd = &cobra.Command{
	Use:     "server",
	Version: "0.0.1",
	Short:   "Start server",
	Long:    "Start server and accept connection from client",
}

const helpTemplate = `{{.Short}}
{{.UsageString}}`

func init() {
	RootCmd.PersistentFlags().IntVarP(
		&Port,
		"port", "p",
		8989,
		"Port for running server",
	)

	RootCmd.PersistentFlags().StringVarP(
		&ServerName,
		"name", "n",
		"chat-server",
		"Chat server name",
	)

}
