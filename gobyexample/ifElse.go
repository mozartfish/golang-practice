package gobyexample

import "fmt"

// Branching with if and else in golang
// note: go does not support ternary statements
func IfElseConstruct() {
	fmt.Println("IF/ELSE")
	// if-else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// if
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// conditional operators
	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// statement conditionals
	// a statement can precede conditionals, any variables declared in this statement are available in the current
	// and all subsequent branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negatove")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
