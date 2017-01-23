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
		// read in input from stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(connection, text+"\n")
		// listen for reply
		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
}
