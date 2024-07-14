package gobyexample

import "fmt"

// For is go's only looping construct.
func ForConstruct() {
	fmt.Println("FOR")
	// while loop
	fmt.Println("while-loop")
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// for loop
	fmt.Println("for-loop")
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	// range loop (for-each)
	fmt.Println("range-loop")
	for i := range 3 {
		fmt.Println("range", i)
	}

	// infinite loop
	fmt.Println("infinite-loop(paused)")
	for {
		fmt.Println("loop")
		break
	}

	// continue statement
	fmt.Println("continue statements")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
