package gomaze

import (
	"fmt"
	"math/rand"
)

type Iterative struct {
	visited map[*Cell]struct{}
	rng     *rand.Rand
}

func NewIterative(seed int64) Initer {
	return &Iterative{
		visited: make(map[*Cell]struct{}),
		rng:     rand.New(rand.NewSource(seed)),
	}
}

func (i *Iterative) Init(g *Grid) error {
	stack := make([]*Cell, 0, 100)

	// Choose a random starting cell
	r := i.rng.Intn(g.Rows())
	c := i.rng.Intn(g.Cols())
	currentCell, err := g.CellAt(c, r)
	if err != nil {
		return fmt.Errorf("init: could not get first cell: %v", err)
	}

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
