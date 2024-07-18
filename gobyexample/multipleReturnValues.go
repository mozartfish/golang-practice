package gobyexample

import "fmt"

// function that returns two values
func vals() (int, int) {
	return 3, 7
}

// Go has builtin support for multiple return values. This featured is used often in idiomatic
// Go ie. to return both result and error values from a function
func MultipleReturnValues() {
	fmt.Println("MULTIPLE RETURN VALUES")
	// store the 2 different return values from the call with multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// rturn subset of the returned values 
	_, c := vals()
	fmt.Println(c)
}
