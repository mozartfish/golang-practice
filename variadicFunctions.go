package main

import "fmt"

// Variadic functions can be called with any number of trailing arguments
// e.g. fmt.Println() is a common variadic function

// function that takes arbitrary number of ints as args and prints total
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func variadicFunctions() {
	fmt.Println("VARIADIC FUNCTIONS")

	// variadic functions are called in the usual way with individual arguments
	sum(1, 2)
	sum(1, 2, 3)

	// if you have multiple args in a slice, apply them to a variadic functions using func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
