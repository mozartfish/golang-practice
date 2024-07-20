package gobyexample

import "fmt"

// define enum type ServerState which has an underlying int type 
type ServerState int

// the possible values for ServerState are defined as constants 
// the special keyword iota generates successive constant values automatically; 
// ie. 0,1, 2, 3, 4, and so on in this example
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// by implementing the fmt.Stringer interface, values of ServerState
// can be printed out or converted to a string
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

// transition emulates a state transition for a serve; it takes the existing state and
// returns a new state
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
	return StateConnected

}

// Enumerate types (enums) are a special case of sum types
// An enum is a type that has a fixed number of possible values, each with a distinct name
// Go doesn't have an enum type as a distinct language feature, but enums are simple to implement
// using existing language idioms
func Enums() {
	fmt.Println("ENUMS")
	// if we have a value of type int, we cannot pass it to transition - the compiler will complain about type mismatch
	// this provides some degree of compile-time type safety for enums
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)
}
