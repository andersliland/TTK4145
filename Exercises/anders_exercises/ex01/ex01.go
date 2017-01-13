// Go 1.2
// go run helloworld_go.go

package main

import (
    "fmt"
    "runtime"
    "time"
)

var Num int  = 0

func someGoroutine() {
    Println("Hello from a goroutine!")
}

func incNumFunc(){
    for i := 0; i < 10; i++ {
        Num++
        Printf("incNum: %v \n", Num) 
    }

}

func decNumFunc(){
    for i := 0; i < 11; i++ {
        Num--
        Printf("decNum: %v \n", Num)
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    go someGoroutine()                      // This spawns someGoroutine() as a goroutine
    go incNumFunc()
    go decNumFunc()

    // We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
    // We'll come back to using channels in Exercise 2. For now: Sleep.
    time.Sleep(100*time.Millisecond)
    Printf("Num: %v" , Num)

}