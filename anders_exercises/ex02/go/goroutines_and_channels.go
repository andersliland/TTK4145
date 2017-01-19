package main

import (
	"fmt"
	"time"
)

func pinger(c chan string){
	for i :=0 ; ; i++ {
		c <- "ping"
	}
}
func printer(c chan string){
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Print("done")

	done <- true
}

func main(){

	done := make(chan bool, 1)
	go worker(done)

	<-done
}