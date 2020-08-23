package cmd

import (
	"github.com/spf13/cobra"
)

var (
	ServerAddr string
	Port       int
	ServerName string
)

var RootCmd = &cobra.Command{
	Use:     "client",
	Version: "0.0.1",
	Short:   "Start Client",
	Long:    "Start client and establish connection to server",
}

const helpTemplate = `{{.Short}}
{{.UsageString}}`

func init() {
	RootCmd.PersistentFlags().IntVarP(
		&Port,
		"port", "p",
		8989,
		"Port of running server",
	)

	RootCmd.PersistentFlags().StringVarP(
		&ServerName,
		"name", "n",
		"chat-client",
		"Chat client name",
	)

	RootCmd.PersistentFlags().StringVarP(
		&ServerAddr,
		"server", "s",
		"127.0.0.1",
		"Server IP to connect",
	)

}
