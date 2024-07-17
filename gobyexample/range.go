package gobyexample

import "fmt"

// range iterates over elements in a variety of data structures
func Range() {
	fmt.Println("RANGE")
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// range on arrays and slices provides both the index and value for each entry
	// the blank identifier _ is used for ignoring either the index or the value
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range over maps iterates over key-value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// rang over map keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points
	// the first value is the starting byte index of the rune
	// the second value is the rune itself
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
