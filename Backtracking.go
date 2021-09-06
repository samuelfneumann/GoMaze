package gomaze

import (
	"math/rand"
)

type Backtracking struct {
	visited map[*Cell]struct{}
	g       *Grid
	rng     *rand.Rand
}

func NewBacktracking(g *Grid, seed int64) Initer {
	return &Backtracking{
		visited: make(map[*Cell]struct{}),
		g:       g,
		rng:     rand.New(rand.NewSource(seed)),
	}
}

func (b *Backtracking) Init() error {
	stack := make([]*Cell, 0, 100)

	// Choose a random starting cell
	r := b.rng.Intn(b.g.Rows())
	c := b.rng.Intn(b.g.Cols())
	currentCell := b.g.CellAt(c, r)

	stack = append(stack, currentCell)
	b.visited[currentCell] = struct{}{}

	for len(stack) > 0 {
		// Choose random unvisited neighbour
		neighbours := make([]*Cell, 0, len(currentCell.Neighbours()))
		for _, cell := range currentCell.Neighbours() {
			if cell == nil {
				continue
			}
			// Check if the cell has been visited
			_, visited := b.visited[cell]
			if !visited {
				neighbours = append(neighbours, cell)
			}
		}

		if len(neighbours) == 0 {
			// If all neighbours have been visited, we are done with
			// the cell and we backtrack to the previously discovered
			// cell
			currentCell = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		} else {
			// An unvisited neighbour was found, move to that cell and
			// mark as the current cell, linking the current cell with
			// its neighbour
			neighbourCell := neighbours[b.rng.Intn(len(neighbours))]
			currentCell.Link(neighbourCell)
			b.visited[neighbourCell] = struct{}{}
			stack = append(stack, neighbourCell)
			currentCell = neighbourCell
		}
	}

	return nil
}
