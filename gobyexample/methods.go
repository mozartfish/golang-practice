package gobyexample

import "fmt"

type rect struct {
	width, height int
}

// pointer receiver
// use pointer receiver to avoid copying on method calls or to allow the method to mutate 
// the receiving struct
func (r *rect) area() int {
	return r.width * r.height
}

// value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// Go supports methods defined on struct types
// Go automatically handles conversion between values and pointers for method calls 
func Methods() {
	fmt.Println("METHODS")
	r := rect{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
