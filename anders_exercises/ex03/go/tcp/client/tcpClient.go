package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	ConnectionHost                    = "localhost"
	ConnectionPortFixesSizeMessage    = "34933"
	ConnectionPort0TerminatedMessafge = "33546"
	ConnectionType                    = "tcp"
)

func main() {

	//connecto to this socket
	conn, err := net.Dail(ConnectionType, ConnectionHost+":"+ConnectionPortFixesSizeMessage)
	if err != nil {
		fmt.Println("Dail error")
	}

	for {
		//read in input from stdin
		reader := bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Print("Text to send: ")

	}
}
