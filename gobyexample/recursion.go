package gobyexample

import "fmt"

// recursive implementation of factorial
func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fact(n-1)
}

// Go supports recursive functions
func Recursion() {
	fmt.Println("RECURSION")

	fmt.Println(" recursion:", fact(7))

	// closures can also be recursive but this requires the closure to be declared
	// with type var explicitly before its defined
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println("closures and recursion:", fib(7))
}
