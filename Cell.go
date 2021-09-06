package gomaze

import (
	"fmt"
	"math/rand"
)

type Cell struct {
	row, col                 int
	north, south, east, west *Cell

	links map[*Cell]struct{} // Can travel to any cell in links
}

func NewCell(r, c int) *Cell {
	return &Cell{
		row:   r,
		col:   c,
		links: make(map[*Cell]struct{}),
	}
}

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

func (c *Cell) Neighbours() []*Cell {
	neighbours := []*Cell{
		c.North(),
		c.South(),
		c.East(),
		c.West(),
	}
	return neighbours
}

func (c *Cell) CanMoveEast() bool {
	_, ok := c.links[c.east]
	return ok
}

func (c *Cell) CanMoveWest() bool {
	_, ok := c.links[c.west]
	return ok
}

func (c *Cell) CanMoveSouth() bool {
	_, ok := c.links[c.south]
	return ok
}

func (c *Cell) CanMoveNorth() bool {
	_, ok := c.links[c.north]
	return ok
}

func (c *Cell) Col() int {
	return c.col
}

func (c *Cell) Row() int {
	return c.row
}

func (c *Cell) North() *Cell {
	return c.north
}

func (c *Cell) East() *Cell {
	return c.east
}

func (c *Cell) West() *Cell {
	return c.west
}

func (c *Cell) South() *Cell {
	return c.south
}

func (c *Cell) Link(new *Cell) {
	c.links[new] = struct{}{}
	new.links[c] = struct{}{}
}

func (c *Cell) Unlink(old *Cell) {
	delete(c.links, old)
	delete(old.links, c)
}

func (c *Cell) Linked(cell *Cell) bool {
	_, ok := c.links[cell]
	return ok
}

func (c *Cell) Links() []*Cell {
	keys := make([]*Cell, len(c.links))
	i := 0
	for key := range c.links {
		keys[i] = key
		i++
	}

	return keys
}
