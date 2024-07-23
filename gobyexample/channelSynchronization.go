package gobyexample

import (
	"fmt"
	"time"
)

// example of using a block receive to wait for a goroutine to finish
// when waiting for multiple goroutines to finish, you may prefer to use a waitgroup
// the done channel will be used to notify another goroutine that this function's work is done 
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value to notify that we're done
	done <- true
}

// we can use channels to synchronize execution across goroutines
func ChannelSynchronization() {
	fmt.Println("CHANNEL SYNCHRONIZATION")

	done := make(chan bool, 1)
	go worker(done)

	// block until we receive notification from the worker on the channel 
	<-done
}
