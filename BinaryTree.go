package gomaze

import (
	"fmt"
	"math/rand"
)

// BinaryTree initializes a grid into a maze using the binary tree
// algorithm.
type BinaryTree struct {
	visited map[*Cell]struct{}
	rng     *rand.Rand
	bias    BiasDirection
}

// NewBinaryTree returns a new BinaryTree
func NewBinaryTree(seed int64) Initer {
	init := &BinaryTree{
		visited: make(map[*Cell]struct{}),
		rng:     rand.New(rand.NewSource(seed)),
	}

	// Set a random bias
	biases := []BiasDirection{NW, NE, SW, SE}

	init.bias = biases[init.rng.Intn(len(biases))]
	return init
}

// NewBinaryTreeWithBias returns a new BinaryTree with bias given by
// bias
func NewBinaryTreeWithBias(seed int64, bias BiasDirection) (Initer, error) {
	if bias != NW && bias != NE && bias != SW && bias != SE {
		return nil, fmt.Errorf("newBinaryTreeWithBias: could not create "+
			"binary tree with unknown bias %v", bias)
	}

	return &BinaryTree{
		visited: make(map[*Cell]struct{}),
		rng:     rand.New(rand.NewSource(seed)),
		bias:    bias,
	}, nil

}

// Init initializes a grid using the binary tree algorithm
func (b *BinaryTree) Init(g *Grid) error {
	f1, f2, err := bias(b.bias)
	if err != nil {
		return fmt.Errorf("init: could not get bias directions: %v", err)
	}

	for r := 0; r < g.Rows(); r++ {
		for c := 0; c < g.Cols(); c++ {
			neighbours := make([]*Cell, 0, 2)
			cell, err := g.CellAt(c, r)
			if err != nil {
				return fmt.Errorf("init: could not get neighbour: %v", err)
			}

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
