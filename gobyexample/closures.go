package gobyexample

import "fmt"

// returns another function defined anonymously in the body of intSeq
// the returned function closes over the variable i to form a closure
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Go supports anonymous functions which can form closures
// anonymous functions are useful when you want to define a function inline without having to name it
func Closures() {
	fmt.Println("CLOSURES")
	// call intSeq assigning the result (a function) to nextInt
	// this function value captures its own i value which will be updated each time we call nextInt
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// the state is unique to a particular function as demoed with newInts
	newInts := intSeq()
	fmt.Println(newInts())
}
