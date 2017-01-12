// Go 1.2
// go run helloworld_go.go

package main

import (
	. "fmt"
	"runtime"
)

func server(ch1, ch2, finished chan bool) int {
	num := 0
	exit := 0
	for {
		select {
		case <-ch1:
			num++
			Println(num)
		case <-ch2:
			num--
			Println(num)
		case <-finished:
			exit++
			if exit == 2 {
				return num
			}
		}
	}
}

// Are all functions Go routines?
func threadIncNum(ch1, finished chan bool) {
	for i := 0; i < 10; i++ {
		ch1 <- true
	}
	finished <- true
}

func threadDecNum(ch2, finished chan bool) {
	for i := 0; i < 10; i++ {
		ch2 <- true
	}
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ch1 := make(chan bool, 1)
	ch2 := make(chan bool, 1)

	finished := make(chan bool, 1)

	go threadIncNum(ch1, finished)
	go threadDecNum(ch2, finished)
	num := server(ch1, ch2, finished)

	Println("Num: ", num)
}
