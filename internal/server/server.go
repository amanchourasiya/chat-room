package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

type Server struct {
	Port int
	Name string
}

// Start will configure and start server
func (server *Server) Start() {
	fmt.Printf("Starting server ...\n")

	port := fmt.Sprintf(":%d", server.Port)
	ln, _ := net.Listen("tcp", port)
	conn, _ := ln.Accept()

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("Message Received: %s\n", string(message))
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter your message: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
	}
}
