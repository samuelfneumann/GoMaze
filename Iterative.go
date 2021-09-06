package gomaze

import (
	"math/rand"
)

type Iterative struct {
	visited map[*Cell]struct{}
	g       *Grid
	rng     *rand.Rand
}

func NewIterative(g *Grid, seed int64) Initer {
	return &Iterative{
		visited: make(map[*Cell]struct{}),
		g:       g,
		rng:     rand.New(rand.NewSource(seed)),
	}
}

func (i *Iterative) Init() error {
	stack := make([]*Cell, 0, 100)

	// Choose a random starting cell
	r := i.rng.Intn(i.g.Rows())
	c := i.rng.Intn(i.g.Cols())
	currentCell := i.g.CellAt(c, r)

	stack = append(stack, currentCell)
	i.visited[currentCell] = struct{}{}

	for len(stack) > 0 {
		// Pop from stack
		currentCell = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Choose random unvisited neighbour
		neighbours := make([]*Cell, 0, len(currentCell.Neighbours()))
		for _, cell := range currentCell.Neighbours() {
			if cell == nil {
				continue
			}
			// Check if the cell has been visited
			_, visited := i.visited[cell]
			if !visited {
				neighbours = append(neighbours, cell)
			}
		}

		if len(neighbours) > 0 {
			stack = append(stack, currentCell)
			neighbourCell := neighbours[i.rng.Intn(len(neighbours))]
			currentCell.Link(neighbourCell)
			i.visited[neighbourCell] = struct{}{}
			stack = append(stack, neighbourCell)
		}
	}

	return nil
}
