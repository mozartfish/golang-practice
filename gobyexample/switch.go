package gobyexample

import (
	"fmt"
	"time"
)

// switch statements express conditionals across many branches
// default is optional for switch statements in golang
func SwitchConstruct() {
	fmt.Println("SWITCH")
	// switch statement
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// swith with multiple expression statements
	// default case is optional
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// swith statement without an expression is an alternate way to express
	// if-else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// type switch
	// type switch compares types instead of values. can be used to discover the type
	// of an interface value ie. the variable t will have the type corresponding to its clause
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	fmt.Println("type switch")
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

}
