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
	// Day 4
	gobyexample.MapConstruct()
	gobyexample.Range()
	gobyexample.Functions()
	// Day 5
	gobyexample.MultipleReturnValues()
	gobyexample.VariadicFunctions()
	gobyexample.Closures()
	// Day 6
	gobyexample.Recursion()
	gobyexample.Pointers()
	gobyexample.StringsRunes()
	// Day 7 
	gobyexample.Structs()
	gobyexample.Methods()
	gobyexample.Interfaces()
	// Day 8 
	gobyexample.Enums()
	gobyexample.StructEmbedding()
	gobyexample.Generics()
	// Day 9 
	gobyexample.Errors()
	gobyexample.CustomErrors()
	gobyexample.Goroutines()
	// Day 10 
	gobyexample.Channels()
	gobyexample.ChannelBuffering()
	gobyexample.ChannelSynchronization()
	// Day 11
	// gobyexample.ChannelDirections()
	// gobyexample.SelectConstruct() 
	// gobyexample.Timeouts()

}
