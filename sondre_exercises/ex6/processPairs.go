package main

import (
	"fmt"
	"time"

	"./UDPbackup"
)

func main() {
	fmt.Println("Backup Mode")
	var counter byte = 1
	listenChan := make(chan byte)
	go UDPbackup.UDPlisten(listenChan)
Backup:
	for {
		select {
		case <-time.After(time.Second):
			break Backup
		case counter = <-listenChan:
			break
		}
	}

	// Spawn backup process

	fmt.Println("Active Mode")
	for {
		fmt.Println(counter)
		counter++
		UDPbackup.UDPsend(counter)
		time.Sleep(500 * time.Millisecond)
	}
}

func launchBackupProcess() {

}
