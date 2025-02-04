// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a custom error type called appError that contains three fields, err error,
// message string and code int. Implement the error interface providing your own message
// using these three fields. Implement a second method named temporary that returns false
// when the value of the code field is 9. Write a function called checkFlag that accepts
// a bool value. If the value is false, return a pointer of your custom error type
// initialized as you like. If the value is true, return a default error. Write a main
// function to call the checkFlag function and check the error using the temporary
// interface.
package main

import (
	"errors"
	"fmt"
)

// Declare a struct type named appError with three fields, err of type error,
// message of type string and code of type int.
type appError struct {
	err     error
	message string
	code    int
}

// Error implements the error interface for appError.
func (e *appError) Error() string {
	return fmt.Sprintf("err=%s, message=%s, code=%d", e.err.Error(), e.message, e.code)
}

// Temporary returns true when the value of the code is not 9.
func (e *appError) Temporary() bool {
	return e.code != 9
}

// Declare the temporary interface type with a method named Temporary that
// takes no parameters and returns a bool.
type temporary interface {
	Temporary() bool
}

// Declare a function named checkFlag that accepts a boolean value and
// returns an error interface value.
func checkFlag(b bool) error {
	// If the parameter is false return an appError.
	if !b {
		return &appError{
			err:     errors.New("app error"),
			message: "message",
			code:    42,
		}
	}

	// Return a default error.
	return errors.New("default error")
}

func main() {
	// Call the checkFlag function to simulate an error of the
	// concrete type.
	err := checkFlag(false)

	// Check the concrete type and handle appropriately.
	switch e := err.(type) {

	// Apply the case for the existence of the Temporary behavior.
	// Log the error and write a second message only if the
	// error is not temporary.
	case temporary:
		fmt.Println(err)
		if !e.Temporary() {
			fmt.Println("error is not temporary")
		}

	// Apply the default case and just log the error.
	default:
		fmt.Println(err)
	}
}
