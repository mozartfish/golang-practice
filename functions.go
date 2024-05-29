package main

import "fmt"

// Functions are central in Go
func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func functionConstruct() {
	fmt.Println("FUNCTIONS")

	res := plus(1, 2)
	fmt.Println("1+2=", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3=", res)
}
