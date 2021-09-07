package gomaze

import (
	"fmt"
	"math/rand"
)

const InitialPathLength int = 10

type Wilson struct {
	free []*Cell
	rng  *rand.Rand

	// To increase efficiency of path initialization
	// The size of the initialized path slice will be biased towards
	// this value unless set to 0
	initialPathLength int
	timesWalkCalled   float64
}

func NewWilson(seed int64) Initer {
	return &Wilson{
		free:              nil,
		rng:               rand.New(rand.NewSource(seed)),
		initialPathLength: InitialPathLength,
	}
}

func (w *Wilson) Init(g *Grid) error {
	w.free = make([]*Cell, g.Len())
	copy(w.free, g.Cells())

	// Get the starting position
	index := w.rng.Intn(g.Len())
	w.free = append(w.free[:index], w.free[index+1:]...)

	for len(w.free) > 0 {
		// Perform a random walk
		path, err := w.walk()
		if err != nil {
			return fmt.Errorf("init: could not walk: %v", err)
		}

		// Update the initial path length with the (biased) estimate of
		// the average path length
		w.initialPathLength += int((1 / w.timesWalkCalled) * float64(
			(len(path) - w.initialPathLength)))

		// Link all cells found in path
		var prev *Cell
		for _, cell := range path {
			if prev != nil {
				_, prevIndex := in(w.free, prev)
				prev.Link(cell)
				w.free = append(w.free[:prevIndex], w.free[prevIndex+1:]...)
			}
			prev = cell
		}
	}
	return nil
}

func (w *Wilson) walk() ([]*Cell, error) {
	w.timesWalkCalled += 1

	// Choose a random starting cell
	index := w.rng.Intn(len(w.free))
	cell := w.free[index]
	if cell == nil {
		panic("cell is nil")
	}

	// Keep track of the path travelled
	path := make([]*Cell, 0, w.initialPathLength)
	path = append(path, cell)

	// Keep find random neighbours until reaching a visited cell
	freeCell, _ := in(w.free, cell)
	for freeCell {
		neighbourCell, err := cell.RandomNeighbour(w.rng)
		if err != nil {
			return nil, fmt.Errorf("walk: could not get neighbour %v", err)
		}

		exists, at := in(path, neighbourCell)
		if exists {
			path = path[:at+1]
		} else {
			path = append(path, neighbourCell)
		}

		cell = neighbourCell
		freeCell, _ = in(w.free, cell)
	}

	return path, nil
}

func in(slice []*Cell, c *Cell) (bool, int) {
	for i, cell := range slice {
		if cell == c {
			return true, i
		}
	}
	return false, -1
}
