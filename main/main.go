package main

import (
	"log"
	"time"

	"github.com/samuelfneumann/gomaze"
)

func main() {
	cols, rows := 15, 10
	startCol, startRow := -1, -1
	goalCol, goalRow := -1, -1

	m, err := gomaze.NewMaze(
		rows,
		cols,
		goalRow,
		goalCol,
		startRow,
		startCol,
		gomaze.NewWilson(time.Now().UnixNano()),
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	m.Play()
}
