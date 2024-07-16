package gobyexample

import "fmt"

// In go an array is a numbered sequence of elements of a specific length
// In typical Go code, slices are much more common; arrays are useful in some special scenarios
func Arrays() {
	fmt.Println("ARRAYS")
	// declare an array that will hold exacly 5 ints. The type of elements and length
	// are both part of the arrays type. By default an array is zero-valued ie. for ints this means 0s
	var a [5]int
	fmt.Println("emp:", a)

	// array get and set methods
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// array length
	fmt.Println("len:", len(a))

	// declare and initialize on 1 line
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// compiler counts the number of elements with ... syntax
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// if you specify index with the : syntax, the elements in between will be zeroed
	b = [...]int{100, 3: 400, 500}
	fmt.Println("idx:", b)

	// multi-dimensional arrays
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	// multi-dimensional arrays declared in 1 line
	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Println("2d: ", twoD)
}
