package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

const (
	host = "10.22.66.249"
	port = "20014"
)

// CheckError ...
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	connection, err := net.Dial("tcp", host+":"+port)
	CheckError(err)
	for {
		// Read input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		CheckError(err)
		// Send to socket
		fmt.Fprintf(connection, text+"\n")
		// Listen for reply
		message, err := bufio.NewReader(connection).ReadString('\n')
		CheckError(err)
		fmt.Print("Message from server: " + message)
	}
}
