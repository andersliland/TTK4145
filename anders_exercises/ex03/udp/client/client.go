package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// CheckError
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	//LocalAddr, err := net.ResolveUDPAddr("udp", "129.241.187.152:")

	//ServerAddr, err := net.ResolveUDPAddr("udp","129.241.187.27:20013")
	ServerAddr, err := net.ResolveUDPAddr("udp", "255.255.255.255:20013")
	CheckError(err)

	Conn, err := net.DialUDP("udp", nil, ServerAddr)
	CheckError(err)
	fmt.Println("DialUDP")
	defer Conn.Close()

	i := 0

	for {
		msg := strconv.Itoa(i)
		i++
		buf := []byte(msg)
		_, err := Conn.Write(buf)
		if err != nil {
			fmt.Println(msg, err)

		}
		time.Sleep(time.Second * 1)

	}
}
