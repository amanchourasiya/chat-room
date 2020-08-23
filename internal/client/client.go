package client

import (
	"fmt"
	"net"
	"sync"

	"github.com/amanchourasiya/chat-room/internal/message"
)

type Client struct {
	ServerAddr string
	Port       int
	Name       string
}

// Start will configure and start client
func (client *Client) Start() {
	conn, _ := net.Dial("tcp", fmt.Sprintf("%s:%d", client.ServerAddr, client.Port))
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
