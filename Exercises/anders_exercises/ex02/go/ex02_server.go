package main

import (
    "fmt"
    "runtime"
)

func server(ch1, ch2, quit chan bool) int {
    var Num int = 0
    var finished int = 0
    for {
        select{
        case <- ch1:
            Num++
            fmt.Println("incNum:", Num)
        case <- ch2:
            Num--
            fmt.Println("decNum:", Num)
        case <- quit:
            finished++
            fmt.Println("incFinished", finished)
            if finished == 2 {
                fmt.Println("quit server:", finished)
                return Num
            }

        }
    }
}


func incNum(ch1, quit chan bool){
    for i := 0; i < 10; i++ {
        ch1 <- true
    }
    quit <- true
}

func decNum(ch2, quit chan bool){
    for i := 0; i < 10; i++ {
        ch2 <- true
    }
    quit <- true
}


func main() {

    ch1 := make(chan bool,10)
    ch2 := make(chan bool,10)
    quit := make(chan bool,10)    

    go incNum(ch1, quit)
    go decNum(ch2, quit)


    fmt.Println("Num:" , server(ch1, ch2 ,quit))

}