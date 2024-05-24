package main

import "fmt"

// Go has various value types including strings, integers floats, booleans, etc.
func values() {
	fmt.Println("VALUES")

	// strings can be added together with +
	fmt.Println("go" + "lang")

	// integers and floats
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// booleans + boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
