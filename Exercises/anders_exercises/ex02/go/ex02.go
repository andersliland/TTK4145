// Go 1.2
// go run helloworld_go.go

package main

import (
    "fmt"
    "runtime"
)

//var Num int  = 0


func incNumFunc(Num int, c chan int){
	var local_num = Num
    for i := 0; i < 10; i++ {
        local_num++
        fmt.Printf("incNum: %v \n", Num) 
    }
    c <- local_num //send local_num to c

}

func decNumFunc(Num int, c chan int){
	var local_num = Num
    for i := 0; i < 11; i++ {
        Num--
        fmt.Printf("decNum: %v \n", Num)
    }
    c <- local_num
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())    // I guess this is a hint to what GOMAXPROCS does...
                                            // Try doing the exercise both with and without it!
    var Num int = 0

    c := make(chan int, 1 )

    go incNumFunc(Num, c)
    go decNumFunc(Num, c)

    
    x := <-c //receiv from c


    fmt.Printf("Num: %v \n" , x)
    //fmt.Printf("Num: %v \n" , y)


}