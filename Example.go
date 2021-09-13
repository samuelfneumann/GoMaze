package gomaze

import (
	"log"
	"time"
)

func Example() {
	cols, rows := 15, 10
	startCol, startRow := -1, -1
	goalCol, goalRow := -1, -1

	m, err := NewMaze(
		rows,
		cols,
		goalRow,
		goalCol,
		startRow,
		startCol,
		NewWilson(time.Now().UnixNano()),
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	m.Play()
}
