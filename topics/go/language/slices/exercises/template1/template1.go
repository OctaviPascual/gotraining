// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import (
	"fmt"
)

func main() {

	// Declare a nil slice of integers.
	var ints []int

	// Append numbers to the slice.
	ints = append(ints, 1, 2, 3)

	// Display each value in the slice.
	for _, v := range ints {
		fmt.Printf("v=%d\n", v)
	}

	// Declare a slice of strings and populate the slice with names.
	strings := []string{"a", "b", "c", "d"}

	// Display each index position and slice value.
	for i, v := range strings {
		fmt.Printf("i=%d, v=%s\n", i, v)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	slice := strings[1:3]

	// Display each index position and slice values for the new slice.
	for i, v := range slice {
		fmt.Printf("i=%d, v=%s\n", i, v)
	}
}
