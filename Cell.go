package gomaze

import (
	"fmt"
	"math/rand"
)

// Cell is a single cell in a grid
type Cell struct {
	row, col                 int   // Cell position
	north, south, east, west *Cell // Neighbour cells

	links map[*Cell]struct{} // Can travel to any cell in links
}

// NewCell creates and returns a new cell at row r and column c
func NewCell(r, c int) *Cell {
	return &Cell{
		row:   r,
		col:   c,
		links: make(map[*Cell]struct{}),
	}
}

// RandomNeighbour returns a random neighbouring cell
func (c *Cell) RandomNeighbour(rng *rand.Rand) (*Cell, error) {
	if c.North() == nil && c.South() == nil && c.East() == nil &&
		c.West() == nil {
		return nil, fmt.Errorf("uniformRandomNeighbour: no non-nil " +
			"neightbour cells")
	}

	var neighbourCell *Cell

	for neighbourCell == nil {
		side := rng.Intn(4)
		switch side {
		case 0:
			neighbourCell = c.north

		case 1:
			neighbourCell = c.south

		case 2:
			neighbourCell = c.east

		case 3:
			neighbourCell = c.west
		}
	}

	return neighbourCell, nil
}

// Neighbours returns the neighbours of c
func (c *Cell) Neighbours() []*Cell {
	neighbours := []*Cell{
		c.North(),
		c.South(),
		c.East(),
		c.West(),
	}
	return neighbours
}

// CanMoveEast returns whether a player can move to the east neighbour
// of the receiver. If there is a wall to the east of the receiver,
// then a player cannot move to the east.
func (c *Cell) CanMoveEast() bool {
	_, ok := c.links[c.east]
	return ok
}

// CanMoveWest returns whether a player can move to the west neighbour
// of the receiver. If there is a wall to the west of the receiver,
// then a player cannot move to the west.
func (c *Cell) CanMoveWest() bool {
	_, ok := c.links[c.west]
	return ok
}

// CanMoveSouth returns whether a player can move to the south neighbour
// of the receiver. If there is a wall to the south of the receiver,
// then a player cannot move to the south
func (c *Cell) CanMoveSouth() bool {
	_, ok := c.links[c.south]
	return ok
}

// CanMoveNorth returns whether a player can move to the north neighbour
// of the receiver. If there is a wall to the north of the receiver,
// then a player cannot move to the north.
func (c *Cell) CanMoveNorth() bool {
	_, ok := c.links[c.north]
	return ok
}

// Col returns the col of the receiver
func (c *Cell) Col() int {
	return c.col
}

// Row returns the row of the receiver
func (c *Cell) Row() int {
	return c.row
}

// North returns the cell to the north of the receiver
func (c *Cell) North() *Cell {
	return c.north
}

// East returns the cell to the east of the receiver
func (c *Cell) East() *Cell {
	return c.east
}

// West returns the cell to the west of the receiver
func (c *Cell) West() *Cell {
	return c.west
}

// South returns the cell to the south of the receiver
func (c *Cell) South() *Cell {
	return c.south
}

// Link links the receiver to new such that a player could move from
// the receiver to new. This is equivalent to removing a wall between
// the receiver and new.
func (c *Cell) Link(new *Cell) {
	c.links[new] = struct{}{}
	new.links[c] = struct{}{}
}

// Unlink unlinks the receiver to old suvh that a player can no longer
// move from the reciver to old. THis is equivalent to adding a wall
// between the reciver ane old.
func (c *Cell) Unlink(old *Cell) {
	delete(c.links, old)
	delete(old.links, c)
}

// Linked returns whether the receiver is linked to cell
func (c *Cell) Linked(cell *Cell) bool {
	_, ok := c.links[cell]
	return ok
}

// Links returns the links of the receiver. Equivalently, this function
// returns the cells that can be moved to from the receiver.
func (c *Cell) Links() []*Cell {
	keys := make([]*Cell, len(c.links))
	i := 0
	for key := range c.links {
		keys[i] = key
		i++
	}

	return keys
}
