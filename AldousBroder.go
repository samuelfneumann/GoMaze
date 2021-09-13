package gomaze

import (
	"fmt"
	"math/rand"
)

// AldousBroder initializes a grid into a maze with the Aldous-Broder
// algorithm.
type AldousBroder struct {
	visited map[*Cell]struct{}
	rng     *rand.Rand
}

// NewAldousBroder returns a new AldousBroder
func NewAldousBroder(seed int64) Initer {
	return &AldousBroder{
		visited: make(map[*Cell]struct{}),
		rng:     rand.New(rand.NewSource(seed)),
	}
}

// Init initialzies a grid into a maze using the AldousBroder algorithm.
func (a *AldousBroder) Init(g *Grid) error {
	// Choose a random starting cell
	r := a.rng.Intn(g.Rows())
	c := a.rng.Intn(g.Cols())
	currentCell, err := g.CellAt(c, r)
	if err != nil {
		return fmt.Errorf("init: could not get first cell: %v", err)
	}

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
