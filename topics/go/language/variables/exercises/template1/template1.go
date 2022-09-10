// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import (
	"fmt"
	"math"
)

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var s1 string
	var i1 int
	var b1 bool

	// Display the value of those variables.
	fmt.Println(s1)
	fmt.Println(i1)
	fmt.Println(b1)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	s2 := "hello"
	i2 := 10
	b2 := true

	// Display the value of those variables.
	fmt.Println(s2)
	fmt.Println(i2)
	fmt.Println(b2)

	// Perform a type conversion.
	var f float32
	f = math.Pi

	// Display the value of that variable.
	fmt.Println(f)
}
