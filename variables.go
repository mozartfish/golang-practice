package main

import "fmt"

// In Go, variables are explicitly declared and used by the compiler to e.g.
// check type-correctness of function calls
func variables() {
	fmt.Println("VARIABLES")

	// declare 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// infer types of variables
	var d = true
	fmt.Println(d)

	// variables declared without a corresponding initialization are zero-valued
	// e.g. zero value for an int is 0
	var e int 
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable
	// e.g. for var f string = "apple" - the syntax is available only inside functions
	f := "apple"
	fmt.Println(f)
}
