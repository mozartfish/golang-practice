package gobyexample

import (
	"fmt"
	"time"
)

func f2(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// A goroutine is a lightweight thread of execution
// when the program is run, the output of the blocking statementis called first 
// then the output of the two goroutines 
// Goroutine output may be interleaved because goroutines are being run concurrently by the go runtime
func Goroutines() {
	fmt.Println("GOROUTINES")

	// call function synchronously
	f2("direct")

	// invoke function as a goroutine using go f(s). This new goroutine will execute concurrently 
	// with the calling one 
	go f2("goroutine")

	// goroutine as an anonymous function call 
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// the two function calls are runing asynchronously in separate goroutines now 
	// wait for them to finish (for a more robust approach use a waitgroup)
	time.Sleep(time.Second)
	fmt.Println("done")
}
