// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a player's batting average. The formula is hits / atBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import (
	"fmt"
)

// Add imports.

// Declare a struct that represents a ball player.
// Include fields called name, atBats and hits.
type player struct {
	name   string
	atBats int
	hits   int
}

// Declare a method that calculates the batting average for a player.
func (p *player) average() float64 {
	if p.atBats == 0 {
		return 0
	}
	return float64(p.hits) / float64(p.atBats)
}

func main() {

	// Create a slice of players and populate each player
	// with field values.
	players := []player{
		{name: "Alice", atBats: 10, hits: 7},
		{name: "Bob", atBats: 12, hits: 6},
		{name: "Charlie"},
	}

	// Display the batting average for each player in the slice.
	for _, player := range players {
		fmt.Printf("%s has a batting average of %.2f\n", player.name, player.average())
	}
}
