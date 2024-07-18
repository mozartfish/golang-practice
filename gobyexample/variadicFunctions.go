package gobyexample

import "fmt"

// function that takes an arbitrary number of ints as arguments 
func sum(nums ...int) {
	// the type of nums is equivalent to []int
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
// Variadic functions can be called with any number of trailing arguments. 
// fmt.Println is a common variadic function 
func VariadicFunctions() {
	fmt.Println("VARIADIC FUNCTIONS")
	sum(1, 2)
	sum (1, 2, 3)

	// if you have multiple args in a slice appluy them to a variadic function using 
	// func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
