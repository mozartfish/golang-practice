package main

import "fmt"

// Compile and run a program in a single file with the same package - go run main.go
// Compile and run a program with multiple files in the same package - go run .
// Build the .go files into an executable and run - go build .
// ./ module_name
func main() {
	fmt.Println("Enter the main function")

	helloWorld()
	values()
	variables()
}
