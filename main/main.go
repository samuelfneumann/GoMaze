package main

import (
	"fmt"
	"time"

	"github.com/samuelfneumann/gomaze"
)

func main() {
	g := gomaze.NewGrid(10, 20)

	w := gomaze.NewAldousBroder(g, time.Now().UnixNano())
	w.Init()

	fmt.Println(g)
}
