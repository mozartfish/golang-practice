package gobyexample

import (
	"fmt"
	"maps"
)

// Maps are Go's builtin associative data type (hashes or dicts in other languages)
// maps print in the form map[k:v k:v]
func MapConstruct() {
	fmt.Println("MAPS")

	// create a new empty map
	m := make(map[string]int)

	// set key-value pairs using name[key] = val
	m["k1"] = 7
	m["k2"] = 13

	// print all key value pairs
	fmt.Println("map:", m)

	// get value associated with a particular key
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if the key does not exist, the zero value of the value type is returned
	// zero value of a map is nil
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// len returns the number of key value pairs when called on a map
	fmt.Println("len:", len(m))

	// delete removes key-value pairs from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// clear removes all key-value pairs from a map
	clear(m)
	fmt.Println("map:", m)

	// the optional second return value when getting a value from a map indicates if
	// the key was present in the map.
	// can be used to disambiguate between missing keys and keys with zero values
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// declare and initialize a new map on 1 line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// utility functions in map package
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

}
