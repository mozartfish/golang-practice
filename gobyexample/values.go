package gobyexample

import "fmt"

// Go has various value types including strings, integers, floats, booleans

func Values() {
	fmt.Println("VALUES")
	// strings, which can be added together with +
	fmt.Println("go" + "lang")

	// integers and floats
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// boolean operators
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
