package main

import "fmt"

// In Go, an array is a numbered sequenee of elements of a specific length.
// In typical Go code, slices are much more common; arrays are useful in some special scenarios

func arrayConstruct() {
	fmt.Println("ARRAYS")

	// create an array that will hold exactly 5 ints
	// the type of elements and length are both part of the arrays type
	// by default an array is zero valued e.g. ints are 0s
	var a [5]int
	fmt.Println("emp:", a)

	// set and get
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	// declare and initialize array in one line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// count the number of elements using the compiler with ...
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// if you specify index with :, the elements in between will be zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// arrays are one dimensional but can be used to compose multi-dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// declare and initialize a multi-dimensional array in one line
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
