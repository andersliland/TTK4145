package main

import "fmt"

func ping(pings chan <- string, msg string) {
	pings <- msg
}

func pong(pings <- chan string, pongs chan <- string) {
	msg := <- pings //receive from channel pings, and assign value to msg
	pongs <- msg // sent msg to channel pongs
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Print(<-pongs) // send value in channel to print
}