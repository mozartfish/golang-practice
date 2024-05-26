package main

import (
	"fmt"
	"math"
)

// Go supports constants of character, string, boolean, and numeric values
const s string = "constant"

func constants() {
	fmt.Println("CONSTANTS")

	// A const statement can appear anywhere a var statement can
	fmt.Println(s)

	// Constant expressions perform arithmetic with arbitrary precision
	const n = 500000000
	const d = 3e20 / n

	fmt.Println(d)

	// A numeric constant has no type until its given one such as by an explicit conversion
	fmt.Println(int64(d))

	// A number can be given a type by using it in a context that requires one
	// such as a variable assignment or function call. Sin expects a float64
	fmt.Println(math.Sin(n))
}
