// https://github.com/gnmathur/aMAZEd
package gomaze

import (
	"strings"
)

type Initer interface {
	Init() error
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
