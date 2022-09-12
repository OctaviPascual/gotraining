// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

import (
	"fmt"
)

// User represents a user
type User struct {
	name  string
	email string
	age   int
}

func main() {
	// Declare variable of type user and init using a struct literal.
	user1 := User{
		name:  "User1",
		email: "user1@gmail.com",
		age:   55,
	}

	// Display the field values.
	fmt.Println(user1.name)
	fmt.Println(user1.email)
	fmt.Println(user1.age)

	// Declare a variable using an anonymous struct.
	user2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "User2",
		email: "user2@gmail.com",
		age:   33,
	}

	// Display the field values.
	fmt.Println(user2.name)
	fmt.Println(user2.email)
	fmt.Println(user2.age)
}
