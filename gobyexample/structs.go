package gobyexample

import "fmt"

// Go's structs are typed collections of fields. They're useful for grouping data together
// to form records.

// the peson struct has name and age fields 
type person struct {
	name string
	age  int
}

// newPerson constructs a new person struct with the given name 
// you can safely return a pointer to local variable as a local variable will survive the
// scope of the function
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func Structs() {
	fmt.Println("STRUCTS")
	// create a new struct 
	fmt.Println(person{"Bob", 20})

	// name fields when initializing a struct 
	fmt.Println(person{name: "Alice", age: 30})

	// omitted values will be zero valued 
	fmt.Println(person{name: "Fred"})

	// &prefix yields a pointer to the struct
	fmt.Println(&person{name: "Ann", age: 40})

	// it's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("Jon"))

	// access a struct with a dot 
	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	// you can use dots with struct pointers - the pointers are automatically dereferenced 
	sp := &s
	fmt.Println(sp.age)

	// structs are mutable 
	sp.age = 51
	fmt.Println(sp.age)

	// if a struct type is only used for a single value, we don't have to give it a name 
	// the value can have an anonymous struct type. this is common for table driven tests 
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

}
