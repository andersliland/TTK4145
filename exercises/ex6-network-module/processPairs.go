package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"./UDPbackup"
)

func main() {
	fmt.Println("Backup Mode")
	var counter byte = 1
	listenChan := make(chan byte)
	quit := make(chan bool)

	flagPtr := flag.Int("flag", 0, "Switch between two ports")
	flag.Parse()
	go UDPbackup.UDPlisten(listenChan, quit, *flagPtr)

Backup:
	for {
		select {
		case <-time.After(time.Second * 2):
			break Backup
		case counter = <-listenChan:
			break
		}
	}

	// Spawn backup process
	close(quit)
	launchBackupProcess(*flagPtr)

	fmt.Println("Active Mode")
	for {
		fmt.Println(counter)
		counter++
		UDPbackup.UDPsend(counter, *flagPtr)
		time.Sleep(500 * time.Millisecond)
	}
}

func launchBackupProcess(flag int) {
	flag = UDPbackup.NotEqual(flag)
	cmd := exec.Command("gnome-terminal", "-x", "sh", "-c", "go run processPairs.go -flag="+strconv.Itoa(flag))
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
