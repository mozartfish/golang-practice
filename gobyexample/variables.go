package gobyexample

import "fmt"

// In Go, variables are explicitly delcared and used by the compiler
// e.g. check type-correctness of function calls

func Variables() {
	fmt.Println("VARIABLES")
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// declare multiple variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go infers the type of initialized variables
	var d = true
	fmt.Println(d)

	// variables declared without a corresponding initialization are zero-valued
	// e.g. int zero-value = 0
	var e int
	fmt.Println(e)

	// the := syntax is shorthand for declaring and initializing a variable
	// the syntax is only available inside functions
	f := "apple"
	fmt.Println(f)

}
