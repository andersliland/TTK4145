package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	broadcast = "255.255.255.255"
	localhost = "127.0.0.1"
	localport = "1234"
	host      = "129.241.187.43"
	port      = "20014"
)

const messageSize = 4 * 1024

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	fmt.Println("Client running")

	ServerAddr, err := net.ResolveUDPAddr("udp", broadcast+":"+port)
	CheckError(err)

	//LocalAddr, err := net.ResolveUDPAddr("udp", localhost+":"+localport)
	//CheckError(err)

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	CheckError(err)

	log.Printf("Established connection to %s \n", ServerAddr)
	log.Printf("Remote UDP adress : %s \n", Conn.RemoteAddr().String())
	log.Printf("Local UDP client adress : %s \n ", Conn.LocalAddr().String())

	defer Conn.Close()

	// Write a message to the server
	i := 0
	for {
		msg := "Sendt to server"
		i++
		writeBuffer := []byte(msg)
		_, err = Conn.Write(writeBuffer)
		fmt.Printf("Send message: %s \n", writeBuffer)
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)

	}

}

// udpTransmittServer handles inncomming connections
func udpTransmittServer() {

}

// udpRecieveServer
func udpRecieveServer(Conn net.Conn) {

	recieveBuffer := make([]byte, messageSize)

	_, err := Conn.Read(recieveBuffer)
	CheckError(err)

}

func udpConnectionReader() {

}
