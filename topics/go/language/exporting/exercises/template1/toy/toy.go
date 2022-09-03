// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Package toy contains support for managing toy inventory.
package toy

// Declare a struct type named Toy with four fields. Name string,
// Weight int, onHand int and sold int.

// Toy is a type that represents a toy
type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

// New returns a new toy
func New(name string, weight int) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

// OnHand returns the current on hand count.
func (t *Toy) OnHand() int {
	return t.onHand
}

// UpdateOnHand updates and returns the current on hand count.
func (t *Toy) UpdateOnHand(count int) int {
	t.onHand += count
	return t.onHand
}

// Sold returns the current sold count.
func (t *Toy) Sold() int {
	return t.sold
}

// SoldCount updates and returns the current sold count.
func (t *Toy) SoldCount(count int) int {
	t.sold += count
	return t.sold
}
