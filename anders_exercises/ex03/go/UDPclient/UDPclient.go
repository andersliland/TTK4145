package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	localhost = "127.0.0.1"
	localport = "0"
	host      = ""
	port      = "30000"
)

func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	fmt.Println("Client running")

	ServerAddr, err := net.ResolveUDPAddr("udp", host+":"+port)
	CheckError(err)

	LocalAddr, err := net.ResolveUDPAddr("udp", localhost+":"+localport)
	CheckError(err)

	Conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	CheckError(err)

	defer Conn.Close()
	i := 0
	for {
		msg := "Sendt to server"
		i++
		buffer := []byte(msg)
		_, err = Conn.Write(buffer)
		fmt.Println("write:", string(buffer))
		if err != nil {
			fmt.Println(msg, err)
		}
		time.Sleep(time.Second * 1)
	}

}
