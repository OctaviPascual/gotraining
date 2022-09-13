// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Create a file with an array of JSON documents that contain a user name and email address. Declare a struct
// type that maps to the JSON document. Using the json package, read the file and create a slice of this struct
// type. Display the slice.
//
// Marshal the slice into pretty print strings and display each element.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Declare a struct type named user with two fields. Name of type string and
// Email of type string. Add tags for each field for the unmarshal call.
type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	// Use the os package to Open the JSON file. Check for errors.
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Could not open file:", err)
		return
	}

	// Schedule the file to be closed once the function returns.
	defer file.Close()

	// Declare a nil slice of user struct types.
	var users []user

	// Decode the JSON from the file into the slice. Check for errors.
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		fmt.Println("Could not decode file:", err)
		return
	}

	// Iterate over the slice and display each user value.
	for _, u := range users {
		fmt.Printf("%+v\n", u)
	}

	// Marshal each user value and display the JSON. Check for errors.
	marshaledUsers, err := json.MarshalIndent(&users, "", "    ")
	if err != nil {
		fmt.Println("Could not marshal users:", err)
		return
	}

	fmt.Println(string(marshaledUsers))
}
