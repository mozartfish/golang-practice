package gobyexample

import (
	"fmt"
	"slices"
)

// Slices are an important data type in Go, giving a more powerful interface to sequences than arrays
func Slices() {
	fmt.Println("SLICES")
	// slices are typed only by the elemnts they contain (not the number of elements)
	// an unintialized slice has value nil (null)
	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	// create an empty slice with nonzero length
	// by default a slice's capacity is equal to its length; if we know the slice is going to grow ahead of time
	// its possible to pass a capacity explicitly as an additional parameter to make
	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// get and set
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// built in slice functions ie. append which returns a new slice containing the value
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// copy a slice into another slice (deep copy)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// slice of slices
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	// initialize a slice on 1 line
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// slices package 
	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	// multi-dimensional slice 
	// the length of the inner slices can vary unlike multi-dimensional arrays
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
