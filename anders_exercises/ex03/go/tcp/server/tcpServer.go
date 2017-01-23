package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// const constants
const (
	ConnectionHost = "localhost"
	ConnectionPort = "3333"
	ConnectionType = "tcp"
)

func main() {
	//Listen for incomming connections
	listener, err := net.Listen(ConnectionType, ConnectionHost+":"+ConnectionPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	conns := clientConns(listener)
	for {
		go handleConn(<-conns)
	}

}

// clientConn
func clientConns(listener net.Listener) chan net.Conn {

	ch := make(chan net.Conn)
	//Close the listener when the application closes
	defer listener.Close()

	i := 0
	go func() {
		for {
			//Accept connection on port
			client, err := listener.Accept()
			if client != nil {
				fmt.Printf("could not accept: " + err.Error())
				continue
			}
			i++
			fmt.Printf("%d: %v <--> %v \n", i, client.LocalAddr(), client.RemoteAddr())
			ch <- client
		}
	}()
	return ch
}

// handleConn
func handleConn(client net.Conn) {
	// Make a buffer to hold incomming data
	buf := make([]byte, 1024)
	b := bufio.NewReader(client)
	//Read the incoming connection into the buffer
	for {
		message, err := b.ReadBytes('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		client.Write(line)
	}
}
