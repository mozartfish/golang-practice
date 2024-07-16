package main

import (
	"fmt"

	// Import gobyexample package
	"github.com/mozartfish/golang-practice/gobyexample"
	// Import gotest package
)

// run programs in the same package
// go run *.go

// build programs into binaries
// go build *.go - * -> represents the file with the main() function
// main is the program entry

func main() {
	fmt.Println("This is the main execution function")
	// Day 1
	gobyexample.HelloWorld()
	gobyexample.Values()
	gobyexample.Variables()
	// Day 2
	gobyexample.Constants()
	gobyexample.ForConstruct()
	gobyexample.IfElseConstruct()
	// Day 3
	gobyexample.SwitchConstruct()
	gobyexample.Arrays()
	gobyexample.Slices()
}
