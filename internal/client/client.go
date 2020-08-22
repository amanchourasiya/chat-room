package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Client struct {
	Port int
	Name string
}

// Start will configure and start client
func (client *Client) Start() {
	conn, _ := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", client.Port))
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter your message: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")

		// waiting for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Server sent: %s\n", message)
	}
}
