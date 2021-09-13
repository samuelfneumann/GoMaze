package gomaze

// Initer initializes a maze from a grid of cells
type Initer interface {
	Init(*Grid) error
}
