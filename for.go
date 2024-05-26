package main

import "fmt"

// for is Go's only looping construct
func forConstruct() {
	fmt.Println("FOR")

	// while loop
	i := 1
	for i <= 3 {
		fmt.Println("while", i)
		i = i + 1
	}

	// for loop
	for j := 0; j < 3; j++ {
		fmt.Println("for", j)
	}

	// for-each loop
	// do this n times -> range over an integer
	for i := range 3 {
		fmt.Println("range", i)
	}

	// infinite loop
	for {
		fmt.Println("infinite", "loop")
		break
	}

	// continue
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
