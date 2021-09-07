package main

import (
	"time"

	"github.com/samuelfneumann/gomaze"
)

func main() {
	m := gomaze.NewMaze(10, 15, gomaze.NewWilson(time.Now().UnixNano()))

	m.Play()
}
