package main

import "fmt"

// range iterates over elements in a variety of data structures
func rangeConstruct() {
	fmt.Println("RANGE")

	// use range to sum the numbers in a slice
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry
	// can ignore the index with the blank identifier _
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key-value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range on map iterate over map keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points
	// The first value is the starting byte index of the rune and the second the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
