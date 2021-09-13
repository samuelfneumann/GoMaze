package gomaze

import (
	"fmt"
	"strings"
)

// Grid implements a grid of cells
type Grid struct {
	rows, cols int
	cells      []*Cell
}

// NewGrid returns a new grid of cells. Each cell has all four walls
// set, so that once in a cell, you cannot move out of the cell.
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

// CellAt returns the cell at column x and row y in the grid
func (g *Grid) CellAt(x, y int) (*Cell, error) {
	if x > g.Cols() {
		return nil, fmt.Errorf("cellAt: column index out of range [%v] with "+
			"length %v", x, g.Cols())
	}
	if y > g.Rows() {
		return nil, fmt.Errorf("cellAt: row index out of range [%v] with "+
			"length %v", y, g.Rows())
	}
	return g.cells[g.Index(x, y)], nil
}

// Index converts an x, y index into the grid into a single-dimensional
// index. If the grid were flattened to be 1-dimensional, the value
// g.Index(x, y) would be the index to cell with column x and row y in
// the grid
func (g *Grid) Index(x, y int) int {
	return y*g.cols + x
}

// Len returns the number of cell in the grid
func (g *Grid) Len() int {
	return g.rows * g.cols
}

// Rows returns the number of row in the grid
func (g *Grid) Rows() int {
	return g.rows
}

// Cols returns the number of columns in the grid
func (g *Grid) Cols() int {
	return g.cols
}

// Cells returns the cells of the grid in a 1-dimensional slice and in
// row-major format.
func (g *Grid) Cells() []*Cell {
	return g.cells
}

// String returns a string representation of the grid
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
