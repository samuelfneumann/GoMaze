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
	grid *Grid
	goal *Cell
	*player
}

func NewMaze(rows, cols int, init Initer) *Maze {
	g := NewGrid(rows, cols)
	init.Init(g)

	return &Maze{
		grid:   g,
		player: newPlayer(g.CellAt(0, 0)),
		goal:   g.CellAt(cols-1, rows-1),
	}
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
	done := m.player.in == m.goal
	if done {
		reward = 0.0
	}

	return obs, reward, done, nil
}

func (m *Maze) Reset() []float64 {
	m.player = newPlayer(m.grid.CellAt(0, 0))

	return []float64{
		float64(m.player.in.Col()),
		float64(m.player.in.Row()),
	}
}

func (m *Maze) String() string {
	var out strings.Builder
	out.WriteString("+")

	for c := 0; c < m.grid.Cols(); c++ {
		out.WriteString("---+")
	}
	out.WriteString("\n")

	for r := 0; r < m.grid.Rows(); r++ {
		top := "|"
		bottom := "+"
		for c := 0; c < m.grid.Cols(); c++ {
			cell := m.grid.cells[m.grid.Index(c, r)]

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
