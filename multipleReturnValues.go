package main

import "fmt"

// Go has builtin support for multiple return values
// This features is used often in idiomatic Go e.g. to return both result and error values from a function

// function that returns 2 ints
func vals() (int, int) {
	return 3, 7
}
func multipleReturnValues() {
	fmt.Println("MULTIPLE RETURN VALUES")

	// return values from function with multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// get subset of values
	_, c := vals()
	fmt.Println(c)
}
