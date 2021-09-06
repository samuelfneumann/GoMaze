// https://github.com/gnmathur/aMAZEd
package gomaze

import (
	"fmt"
	"math/rand"
	"strings"
)

type Initer interface {
	Init() error
}

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

type Grid struct {
	rows, cols int
	cells      []*Cell
}

func NewGrid(rows, cols int) *Grid {
	cells := make([]*Cell, rows*cols)

	g := &Grid{
		rows: rows,
		cols: cols,
	}

	// Create the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			cells[g.Index(c, r)] = NewCell(r, c)
		}
	}

	// Set neighbouring cells
	for i := range cells {
		row, col := cells[i].Row(), cells[i].Col()
		cell := cells[i]

		if row < rows && row > 0 {
			cell.north = cells[g.Index(col, row-1)]
		}

		if row < rows-1 && row >= 0 {
			cell.south = cells[g.Index(col, row+1)]
		}

		if col < cols && col > 0 {
			cell.west = cells[g.Index(col-1, row)]
		}

		if col < cols-1 && col >= 0 {
			cell.east = cells[g.Index(col+1, row)]
		}
	}

	g.cells = cells
	return g
}

func (g *Grid) CellAt(x, y int) *Cell {
	return g.cells[g.Index(x, y)]
}

func (g *Grid) Index(x, y int) int {
	return y*g.cols + x
}

func (g *Grid) Len() int {
	return g.rows * g.cols
}

func (g *Grid) Rows() int {
	return g.rows
}

func (g *Grid) Cols() int {
	return g.cols
}

func (g *Grid) Cells() []*Cell {
	return g.cells
}

func (g *Grid) String() string {
	var out strings.Builder
	out.WriteString("+")

	for c := 0; c < g.Cols(); c++ {
		out.WriteString("---+")
	}
	out.WriteString("\n")

	for r := 0; r < g.Rows(); r++ {
		top := "|"
		bottom := "+"
		for c := 0; c < g.Cols(); c++ {
			cell := g.cells[g.Index(c, r)]
			body := "   "
			var eastBoundary string
			if cell.Linked(cell.East()) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}
			top = top + body + eastBoundary

			var southBoundary string
			if cell.Linked(cell.South()) {
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}
			corner := "+"
			bottom = bottom + southBoundary + corner
		}
		out.WriteString(top)
		out.WriteString("\n")
		out.WriteString(bottom)
		out.WriteString("\n")
	}
	return out.String()
}
