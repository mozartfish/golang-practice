package gobyexample

import (
	"errors"
	"fmt"
)

// A custom error type usually has the suffix error
type argError struct {
	arg     int
	message string
}

// adding this error method makes argError implement the error interfae
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		// return custom error
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

// It's possible to use custom types as errors by implementing the Error() method on them
// Here's a variant on the example above that uses a custom type to explicitly represent an argument error
func CustomErrors() {
	fmt.Println("CUSTOM ERRORS")
	_, err := f1(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
