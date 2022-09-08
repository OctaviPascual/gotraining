// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Write a program that creates a fixed set of workers to generate random
// numbers. Discard any number divisible by 2. Continue receiving until 100
// numbers are received. Tell the workers to shut down before terminating.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	// Create the channel for sharing results.
	ch := make(chan int)

	// Create a channel "shutdown" to tell goroutines when to terminate.
	shutdown := make(chan struct{})

	// Define the size of the worker pool. Use runtime.GOMAXPROCS(0) to size the pool based on number of processors.
	workers := runtime.GOMAXPROCS(0)

	// Create a sync.WaitGroup to monitor the Goroutine pool. Add the count.
	var wg sync.WaitGroup
	wg.Add(workers)

	// Create a fixed size pool of goroutines to generate random numbers.
	for i := 0; i < workers; i++ {
		go func() {

			// Start an infinite loop.
			for {

				// Generate a random number up to 1000.
				n := rand.Intn(1000)

				// Use a select to either send the number or receive the shutdown signal.
				select {
				// In one case send the random number.
				case ch <- n:
				// In another case receive from the shutdown channel.
				case <-shutdown:
					wg.Done()
					return
				}
			}
		}()
	}

	// Create a slice to hold the random numbers.
	values := make([]int, 0, 100)

	// Receive from the values channel with range.
	for value := range ch {

		// continue the loop if the value was even.
		if value%2 == 0 {
			continue
		}

		// Store the odd number.
		values = append(values, value)

		// break the loop once we have 100 results.
		if len(values) == 100 {
			break
		}
	}

	// Send the shutdown signal by closing the shutdown channel.
	close(shutdown)

	// Wait for the Goroutines to finish.
	wg.Wait()

	// Print the values in our slice.
	fmt.Println(values)
}
