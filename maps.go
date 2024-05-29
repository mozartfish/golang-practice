package main

import (
	"fmt"
	"maps"
)

// Maps are Go's builtin associative data type (hashes/dicts in other languages)

func mapConstruct() {
	fmt.Println("MAPS")

	// create an empty map using the builtin make
	m := make(map[string]int)

	// set key value pairs
	m["k1"] = 7
	m["k2"] = 13
	// print key-value pairs in the format: key-value
	fmt.Println("map:", m)

	// get value
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if key doesn't exist, the zero value of the value type is returned
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	// remove key value pair from a map
	delete(m, "k2")
	fmt.Println("map:", m)

	// remove all key value pairs from a map
	clear(m)
	fmt.Println("map:", m)

	// the optional second return value when getting a value from a map indicates if
	// the key was present in the map. Can be used to disambiguate between missing keys
	// and keys with zero values like 0 or "".
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// initialize a new map in a single line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n === n2")
	}

}
