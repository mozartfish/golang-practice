package gobyexample

import (
	"fmt"
	"math"
)

// interface for geometry shapes 
type geometry interface {
	area() float64
	perim() float64
}

// implement geometry interface on rect and circle types 
type rect1 struct {
	width, height float64
}

type circle1 struct {
	radius float64
}

// to implement an interface in Go, we need to implement all the methods in the interface 
// implement geometry on rects
func (r rect1) area() float64 {
	return r.width * r.height
}

func (r rect1) perim() float64 {
	return 2*r.width + 2*r.height
}

// implement geometry on circles
func (c circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle1) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type, then we can call methods that are in the named interface
// this is a generic measure function taking advantage of this to work on any geometry
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func Interfaces() {
	fmt.Println("INTERFACES")

	r := rect1{width: 3, height: 4}
	c := circle1{radius: 5}

	// rect and circle structs both implement the geometry interface so we can use instances of these 
	// structs as arguments to measure
	measure(r)
	measure(c)
}
