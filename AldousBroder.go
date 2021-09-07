package gomaze

import (
	"fmt"
	"math/rand"
)

type AldousBroder struct {
	visited map[*Cell]struct{}
	rng     *rand.Rand
}

func NewAldousBroder(seed int64) Initer {
	return &AldousBroder{
		visited: make(map[*Cell]struct{}),
		rng:     rand.New(rand.NewSource(seed)),
	}
}

func (a *AldousBroder) Init(g *Grid) error {
	// Choose a random starting cell
	r := a.rng.Intn(g.Rows())
	c := a.rng.Intn(g.Cols())
	currentCell := g.CellAt(c, r)

	a.visited[currentCell] = struct{}{}
	numVisited := 1

	for numVisited < g.Len() {
		neighbourCell, err := currentCell.RandomNeighbour(a.rng)
		if err != nil {
			return fmt.Errorf("init: could not get neighbour: %v", err)
		}

		_, visited := a.visited[neighbourCell]
		if !visited {
			currentCell.Link(neighbourCell)
			a.visited[neighbourCell] = struct{}{}
			numVisited++
		}
		currentCell = neighbourCell
	}

	return nil
}
