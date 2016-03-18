package main

import (
	"fmt"
)

/*
	Channels are the pipes that connect concurrent goroutines. 
	You can send values into channels from one goroutine and receive those values into another goroutine.
*/
func main() {

	/*
		Create a new channel with make(chan val-type). 
		Channels are typed by the values they convey.
		(Tipo de valores que transmiten)
	*/
	messages := make(chan string)
	/*
		Send a value into a channel using the channel <- syntax. 
		Here we send "ping" to the messages channel we made above, from a new goroutine.
	*/
	go func() { messages <- "ping" }()

	/*
		The <-channel syntax receives a value from the channel. 
		Here we’ll receive the "ping" message we sent above and print it out.
	*/
	msg := <-messages
    fmt.Println(msg)
	
}