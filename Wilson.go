package gomaze

import "math/rand"

type Wilson struct {
	free []*Cell
	g    *Grid
	rng  *rand.Rand
}

func NewWilson(g *Grid, seed int64) Initer {
	free := make([]*Cell, g.Len())
	copy(free, g.cells)

	return &Wilson{
		free: free,
		g:    g,
		rng:  rand.New(rand.NewSource(seed)),
	}
}

func (w *Wilson) Init() error {
	// Get the starting position
	index := w.rng.Intn(w.g.Len())
	w.free = append(w.free[:index], w.free[index+1:]...)

	for len(w.free) > 0 {
		// Perform a random walk
		path := w.walk()

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

func (w *Wilson) walk() []*Cell {
	// Choose a random starting cell
	index := w.rng.Intn(len(w.free))
	cell := w.free[index]
	if cell == nil {
		panic("cell is nil")
	}

	// Keep track of the path travelled
	path := make([]*Cell, 0, 10)
	path = append(path, cell)

	// Keep find random neighbours until reaching a visited cell
	freeCell, _ := in(w.free, cell)
	for freeCell {
		var neighbourCell *Cell

		for neighbourCell == nil {
			side := w.rng.Intn(4)
			switch side {
			case 0:
				neighbourCell = cell.north

			case 1:
				neighbourCell = cell.south

			case 2:
				neighbourCell = cell.east

			case 3:
				neighbourCell = cell.west
			}
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
	return path
}

func in(slice []*Cell, c *Cell) (bool, int) {
	for i, cell := range slice {
		if cell == c {
			return true, i
		}
	}
	return false, -1
}
