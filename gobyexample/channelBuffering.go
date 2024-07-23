package gobyexample

import "fmt"

// by default channels are unbuffered, ,eaning that they will only accept sends (chan <-)
// if there is a corresponding receive (<-chan) ready to receive the sent value
// buffered channels accept a limited number of values without a corresponding receiver for those values
func ChannelBuffering() {
	fmt.Println("CHANNEL BUFFERING")

	// channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// because this channel is buffered, we can send these vaoues into the channel without 
	// a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	// receive and print values
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
