package gobyexample

import "fmt"

// takes two ints and returns their sum as an int
func plus(a int, b int) int {
	// Go requires explicit returns ie. it won't automatically return the value of the last expression
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// functions are central in Go
func Functions() {
	fmt.Println("FUNCTIONS")
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
