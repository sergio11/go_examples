package main

import (
	"fmt"
)

//you can specify if a channel is meant to only send or receive values. 


/*
	This ping function only accepts a channel for sending values. 
	It would be a compile-time error to try to receive on this channel.
*/

// chan<- (Sólo enviar) , <-chan (Sólo recibir)
func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

	
func main() {
	/*
		When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
		This specificity increases the type-safety of the program.
	*/
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
