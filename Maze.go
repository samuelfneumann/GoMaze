package gomaze

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Actions = 4 // Number of actions

type player struct {
	in *Cell
}

func newPlayer(start *Cell) *player {
	return &player{
		in: start,
	}
}

func (p *player) MoveSouth() {
	if p.in.CanMoveSouth() {
		p.in = p.in.South()
	}
}

func (p *player) MoveNorth() {
	if p.in.CanMoveNorth() {
		p.in = p.in.North()
	}
}

func (p *player) MoveWest() {
	if p.in.CanMoveWest() {
		p.in = p.in.West()
	}
}

func (p *player) MoveEast() {
	if p.in.CanMoveEast() {
		p.in = p.in.East()
	}
}

type Maze struct {
	*Grid
	goal  *Cell
	start *Cell
	*player
}

// NewMaze returns a new maze of dimensions rows ⨉ cols. The goal
// position is at (goalCol, goalRow). If goalCol or goalRow is less
// than 0, then the bottom right cell is used as the goal. The starting
// position is at (startCol, startRow). If startCol or startRow is
// less than 0, thyen the top left cell is used as the starting
// cell.
func NewMaze(rows, cols int, goalRow, goalCol int,
	startRow, startCol int, init Initer) (*Maze, error) {
	g := NewGrid(rows, cols)
	init.Init(g)

	// Get the goal cell
	var goal *Cell
	var err error
	if goalRow < 0 || goalCol < 0 {
		goal, err = g.CellAt(cols-1, rows-1)
	} else {
		goal, err = g.CellAt(goalRow, goalCol)
	}
	if err != nil {
		return nil, fmt.Errorf("newMaze: could not get goal position: %v",
			err)
	}

	// Get the starting cell
	var playerStart *Cell
	if startRow < 0 || startCol < 0 {
		playerStart, err = g.CellAt(0, 0)
	} else {
		playerStart, err = g.CellAt(startRow, startCol)
	}
	if err != nil {
		return nil, fmt.Errorf("newMaze: could not get start position: %v",
			err)
	}

	return &Maze{
		Grid:   g,
		player: newPlayer(playerStart),
		goal:   goal,
		start:  playerStart,
	}, nil
}

// SetCell sets the current cell of the player
func (m *Maze) SetCell(col, row int) error {
	cell, err := m.CellAt(col, row)
	if err != nil {
		return fmt.Errorf("setCell: %v", err)
	}

	m.player.in = cell

	return nil
}

func (m *Maze) Goal() (int, int) {
	return m.goal.Row(), m.goal.Col()
}

func (m *Maze) AtGoal() bool {
	return m.player.in == m.goal
}

func (m *Maze) Step(action int) ([]float64, float64, bool, error) {
	if action < 0 || action > Actions {
		return nil, 0, false, fmt.Errorf("step: invalid action %v ∉ [%v, %v)",
			action, 0, Actions)
	}

	switch action {
	case 0:
		m.MoveNorth()

	case 1:
		m.MoveSouth()

	case 2:
		m.MoveWest()

	case 3:
		m.MoveEast()
	}

	obs := []float64{
		float64(m.player.in.Col()),
		float64(m.player.in.Row()),
	}
	reward := -1.0
	done := m.AtGoal()
	if done {
		reward = 0.0
	}

	return obs, reward, done, nil
}

func (m *Maze) Reset() []float64 {
	m.player = newPlayer(m.start)

	return []float64{
		float64(m.player.in.Col()),
		float64(m.player.in.Row()),
	}
}

func (m *Maze) String() string {
	var out strings.Builder
	out.WriteString("+")

	for c := 0; c < m.Cols(); c++ {
		out.WriteString("---+")
	}
	out.WriteString("\n")

	for r := 0; r < m.Rows(); r++ {
		top := "|"
		bottom := "+"
		for c := 0; c < m.Cols(); c++ {
			cell := m.cells[m.Index(c, r)]

			var body string
			if cell == m.player.in {
				body = " x "
			} else {
				body = "   "
			}

			if cell == m.goal {
				body = " 🏳 "
			} else if cell != m.player.in {
				body = "   "
			}

			var eastBoundary string
			if cell.Linked(cell.East()) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}
			top = top + body + eastBoundary

			var southBoundary string
			if cell.Linked(cell.South()) {
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}
			corner := "+"
			bottom = bottom + southBoundary + corner
		}
		out.WriteString(top)
		out.WriteString("\n")
		out.WriteString(bottom)
		out.WriteString("\n")
	}
	return out.String()
}

// Play runs the maze game in an interactive session
func (m *Maze) Play() {
	reader := bufio.NewReader(os.Stdin)

	for m.player.in != m.goal {
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println(m)
		fmt.Printf("Action [W S A D; Q - Quit]: ")

		line, err := reader.ReadString('\n')
		if err != nil {
			panic(fmt.Sprintf("play: could not read input: %v", err))
		}

		switch strings.ToUpper(line)[0] {
		case 'W':
			m.MoveNorth()

		case 'S':
			m.MoveSouth()

		case 'A':
			m.MoveWest()

		case 'D':
			m.MoveEast()

		case 'Q':
			os.Exit(0)

		default:
			reader.Reset(os.Stdin)
			fmt.Printf("\rAction [W S A D; Q - Quit]: ")
		}
	}

	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(m)
	fmt.Println("You won!")
}
