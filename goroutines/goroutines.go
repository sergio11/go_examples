package main

import (
	"fmt"
	"time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {

	//running it synchronously.
	f("direct")
	//A goroutine is a lightweight thread of execution.
	go f("goroutine")
	go f("goroutine 2")
	go f("goroutine 3")
	//start a goroutine for an anonymous function call.
	go func(msg string) {
        fmt.Println(msg)
    }("anonymous function call")

    


    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}