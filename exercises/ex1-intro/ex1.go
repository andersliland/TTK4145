// Go 1.2
// go run helloworld_go.go

package main

import (
	"fmt"
	"runtime"
	"time"
)

var num int

// Are all functions go routines?
func threadIncNum() {
	for i := 0; i < 1000000; i++ {
		num++
	}
}

func threadDecNum() {
	for i := 0; i < 1000000; i++ {
		num--
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // I guess this is a hint to what GOMAXPROCS does...
	// Try doing the exercise both with and without it!
	go threadIncNum() // This spawns threadIncNum() as a goroutine
	go threadDecNum()
	// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
	// We'll come back to using channels in Exercise 2. For now: Sleep.
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Num: ", num)
}
