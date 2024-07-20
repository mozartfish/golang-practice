package gobyexample

import "fmt"

// mapkeys takes a map of any type and returns a slice of its keys
// mapkeys has to type parameters - K and V
// K -> has the comparable constraint meaning that we can compare values of this type with the == and != operators
// which is required for map keys in Go
// V -> has any constraint, meaning that it's not restricted in any way (any is an alias for interface{})
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

// generic linked list
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// linked list methods
// we can define methods on generic types just like we do regular types but we have to keep the
// type parameters in place. The type is List[T], not List
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// starting go 1.18, GO has added support for generics also known as type parameters
// note: the order of iteration over map keys is not defined in Go, so different invocations
// may result in different orders
func Generics() {
	fmt.Println("GENERICS")
	// when invoking generic functions, we can often rely on type inference
	// note that we don't have to specify the types for K and V when calling MapKeys - the compiler infers them automatically
	// though we could also specify them explicitly
	var m = map[int]string{1: "2", 2: "4", 4: "8"}

	fmt.Println("keys:", MapKeys(m))

	_ = MapKeys[int, string](m)
	lst := List[int]{}
	lst.Push(10)
	lst.Push(13)
	lst.Push(23)
	fmt.Println("list:", lst.GetAll())

}
