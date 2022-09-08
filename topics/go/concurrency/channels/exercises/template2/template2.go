// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that uses a fan out pattern to generate 100 random numbers
// concurrently. Have each goroutine generate a single random number and return
// that number to the main goroutine over a buffered channel. Set the size of
// the buffered channel so no send ever blocks. Don't allocate more capacity
// than you need. Have the main goroutine store each random number it receives
// in a slice. Print the slice values then terminate the program.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Declare constant for number of goroutines.
const goroutines = 10

func init() {
	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the buffered channel with room for
	// each goroutine to be created.
	ch := make(chan int, goroutines)

	// Iterate and launch each goroutine.
	for i := 0; i < goroutines; i++ {

		// Create an anonymous function for each goroutine that
		// generates a random number and sends it on the channel.
		go func() { ch <- rand.Intn(10) }()
	}

	// Create a variable to be used to track received messages.
	// Set the value to the number of goroutines created.
	received := goroutines

	// Iterate receiving each value until they are all received.
	// Store them in a slice of ints.
	values := make([]int, 0, goroutines)
	for received > 0 {
		value := <-ch
		values = append(values, value)
		received--
	}

	// Print the values in our slice.
	fmt.Println(values)
}
