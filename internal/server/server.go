package server

import (
	"fmt"
	"net"
	"sync"

	"github.com/amanchourasiya/chat-room/internal/message"
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

	sendChan := make(chan string)
	receiveChan := make(chan string)
	var wg sync.WaitGroup
	wg.Add(4)

	fmt.Printf("Starting all go routines\n")
	go message.Receive(&wg, receiveChan, conn)
	go message.PrintMessage(&wg, receiveChan)
	go message.Send(&wg, sendChan, conn)
	go message.ReadMessage(&wg, sendChan)
	fmt.Printf("Finished all go routines\n")
	wg.Wait()
}
