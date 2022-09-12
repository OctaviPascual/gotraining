// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that inspects a user's name and greets them in a certain way
// if they are on a list or in a different way if they are not. Also look at
// the user's age and tell them some special secret if they are old enough to
// know it.
package main

import (
	"fmt"
)

func main() {

	// Change these values and rerun your program.
	name := "Carter"
	age := 6

	// If the user's name is on a special list then give them a secret greeting.
	special := false
	for _, n := range []string{"Alice", "Bob"} {
		if name == n {
			special = true
		}
	}

	if special {
		fmt.Println("You are special!")
	} else {
		fmt.Println("Welcome!")
	}

	// If the user is old enough then tell them a secret.
	if age > 18 {
		fmt.Println("You are of legal age")
	}
}
