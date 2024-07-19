package gobyexample

import "fmt"

// Pass By Value
// zeroval has an int parameter so arguments will be passed to it by value
// zeroval will get a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 0
}

// Pass By Reference
// zeroptr has an *int parameter meaning that it takes an int pointer
// the *iptr code in the function body then dereferences the pointers from its memory address
// to the current value at that address. Assigning a value to a dereferened pointer changes
// the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

// Go supports pointers, allowing you to pass references to values and records within your program
func Pointers() {
	// * in a type means the address of a 
	// * as an operator means value at the address 
	// & as an operator means the address of 
	fmt.Println("POINTERS")
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// &i gives the memory address of i ie. a pointer to i 
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Print pointers
	fmt.Println("pointer:", &i)
}
