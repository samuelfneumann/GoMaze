package main

import (
	"time"

	"github.com/samuelfneumann/gomaze"
)

func main() {
	m := gomaze.NewMaze(5, 5, gomaze.NewWilson(time.Now().UnixNano()))

	m.Play()
}
