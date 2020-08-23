package message

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

func Receive(wg *sync.WaitGroup, ch chan<- string, conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Closing channel as nothing to read %v\n", err)
			close(ch)
			break
		}
		ch <- message
	}
	wg.Done()
}

func Send(wg *sync.WaitGroup, ch <-chan string, conn net.Conn) {
	for {
		if message, ok := <-ch; ok {
			fmt.Fprintf(conn, message+"\n")
		} else {
			fmt.Printf("Connection closed\n")
			break
		}
	}
	wg.Done()

}

func PrintMessage(wg *sync.WaitGroup, ch <-chan string) {
	for {
		if message, ok := <-ch; ok {
			fmt.Printf("Server sent: %s\n", message)
		} else {
			fmt.Printf("Connection closed\n")
			break
		}
	}
	wg.Done()
}

func ReadMessage(wg *sync.WaitGroup, ch chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error in receiving message\n")
			break
		}
		ch <- message
	}
	wg.Done()
}
