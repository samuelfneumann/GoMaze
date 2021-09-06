package gomaze

import (
	"fmt"
	"math/rand"
)

type AldousBroder struct {
	visited map[*Cell]struct{}
	g       *Grid
	rng     *rand.Rand
}

func NewAldousBroder(g *Grid, seed int64) Initer {
	return &AldousBroder{
		visited: make(map[*Cell]struct{}),
		g:       g,
		rng:     rand.New(rand.NewSource(seed)),
	}
}

func (a *AldousBroder) Init() error {
	// Choose a random starting cell
	r := a.rng.Intn(a.g.Rows())
	c := a.rng.Intn(a.g.Cols())
	currentCell := a.g.CellAt(c, r)

	a.visited[currentCell] = struct{}{}
	numVisited := 1

	for numVisited < a.g.Len() {
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
