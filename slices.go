package main

import (
	"fmt"
	"slices"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays
func sliceConstruct() {
	fmt.Println("SLICES")

	// slices are typed only by the elements they contain (not the number of elements)
	// an uninitialized slice euqals to nil and has length 0 
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// to create an empty slice with nonzero slice use builtin make 
	// e.g. create a slice of string of length 3(initially zero valued)
	// by default a new slices capacity is equal to its length; if we know the slice is going to grow ahead of time 
	// it's possible to pass a capacity explicitly as an additional parameter
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// get and set 
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	// slice append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// slice copy 
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slices of slices -> slices of slices are shallow copies and manipualte the original array 
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// initialize and declare a slice in a single line 
	t2 := []string{"g", "h", "i"}
	fmt.Println("dcl:", t2)

	// slice utility functions
	t3 := []string{"g", "h", "i"}
	if slices.Equal(t2, t3) {
		fmt.Println("t2 === t3")
	}
	
	// slices can be used to compose multidimensional slices 
	twoD2 := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD2[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD2[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD2)

}
