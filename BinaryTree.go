package gomaze

import (
	"fmt"
	"math/rand"
)

type BinaryTree struct {
	visited map[*Cell]struct{}
	g       *Grid
	rng     *rand.Rand
	bias    BiasDirection
}

func NewBinaryTree(g *Grid, seed int64, bias BiasDirection) Initer {
	return &BinaryTree{
		visited: make(map[*Cell]struct{}),
		g:       g,
		rng:     rand.New(rand.NewSource(seed)),
		bias:    bias,
	}
}

func (b *BinaryTree) Init() error {
	f1, f2, err := bias(b.bias)
	if err != nil {
		return fmt.Errorf("init: could not get bias directions: %v", err)
	}

	for r := 0; r < b.g.Rows(); r++ {
		for c := 0; c < b.g.Cols(); c++ {
			neighbours := make([]*Cell, 0, 2)
			cell := b.g.CellAt(c, r)

			if f1(cell) != nil {
				neighbours = append(neighbours, f1(cell))
			}
			if f2(cell) != nil {
				neighbours = append(neighbours, f2(cell))
			}

			if len(neighbours) > 0 {
				index := b.rng.Intn(len(neighbours))
				neighbourCell := neighbours[index]
				cell.Link(neighbourCell)
			}
		}
	}
	return nil
}

type BiasDirection int

const (
	NW BiasDirection = iota
	NE
	SW
	SE
)

func bias(direction BiasDirection) (func(c *Cell) *Cell, func(c *Cell) *Cell,
	error) {
	switch direction {
	case NW:
		return (*Cell).North, (*Cell).West, nil

	case NE:
		return (*Cell).North, (*Cell).East, nil

	case SW:
		return (*Cell).South, (*Cell).West, nil

	case SE:
		return (*Cell).South, (*Cell).East, nil

	default:
		return nil, nil, fmt.Errorf("bias: no such bias: %v", direction)
	}
}
