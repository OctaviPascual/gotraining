// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program where two goroutines pass an integer back and forth
// ten times. Display when each goroutine receives the integer. Increment
// the integer with each pass. Once the integer equals ten, terminate
// the program cleanly.
package main

import (
	"fmt"
	"sync"
)

func main() {

	// Create an unbuffered channel.
	ch := make(chan int)

	// Create the WaitGroup and add a count
	// of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("A", ch)
		wg.Done()
	}()

	// Launch the goroutine and handle Done.
	go func() {
		goroutine("B", ch)
		wg.Done()
	}()

	// Send a value to start the counting.
	ch <- 1

	// Wait for the program to finish.
	wg.Wait()
}

// goroutine simulates sharing a value.
func goroutine(id string, ch chan int) {
	for {
		// Wait for the value to be sent.
		// If the channel was closed, return.
		value, ok := <-ch
		if !ok {
			fmt.Printf("[%s] channel is closed, returning\n", id)
			return
		}

		// Display the value.
		fmt.Printf("[%s] received value: %d\n", id, value)

		// Terminate when the value is 10.
		if value == 10 {
			close(ch)
			fmt.Printf("[%s] value is 10, closing channel and returning\n", id)
			return
		}

		// Increment the value and send it
		// over the channel.
		ch <- value + 1
	}
}
