package main

import "fmt"

// Go supports anonymous functions which can form closures
// Anonymous functions are useful when you want to define a function inline without having to name it

// returns another function which we define anonymously in the function body
// the returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func closures() {
	fmt.Println("CLOSURES")

	// assign the result(a function) to next int
	// this value captures its own i value, which will be updated each time nextInt is called
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// confirm that the state is unique to that particular function
	newInts := intSeq()
	fmt.Println(newInts())
}
