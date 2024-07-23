package gobyexample

import "fmt"

// Channels are the pipes that connect concurrent goroutines
// You can send values into channels from one goroutine and receive those values into another goroutine
func Channels() {
	fmt.Println("CHANNELS")
	// create a new channel with make -> channels are typed by the values they convey
	messages := make(chan string)

	// send a value into a channel using the channel <- syntax
	// here ping is sent to the messages channel from a new goroutine
	go func() {
		messages <- "ping"
	}()

	// the <- channel syntax receives a value from the channel
	// here ping is received and is then printed
	// by default sends and receives block UNTIL both the sender and receiver are ready
	// this property allows us to wait at the end of the program for the ping message without
	// having to use any other synchronization
	msg := <-messages
	fmt.Println(msg)
}
