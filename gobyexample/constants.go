package gobyexample

import (
	"fmt"
	"math"
)

// const declares a constant value
const s string = "constant"

// Go supports constants of character, string, boolean and numeric values
func Constants() {
	fmt.Println("CONSTANTS")
	fmt.Println(s)

	// a const statement can appear anywhere a var statement can
	const n = 500000000

	// constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	// a numeric constant has no type until its given one ie. explicit
	fmt.Println(int64(d))

	// a number can be given a type by using it in a context that requires one ie.
	// ie. a single variable assignment or function call
	fmt.Println(math.Sin(n))
}
