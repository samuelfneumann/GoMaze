package gomaze

type Initer interface {
	Init(*Grid) error
}

// // Init initializes g with a copy of init.
// func Init(g *Grid, seed int64, init Initer) error {
// 	switch init.(type) {

// 	case *AldousBroder:
// 		init = NewAldousBroder(g, seed)

// 	case *Backtracking:
// 		init = NewBacktracking(g, seed)

// 	case *BinaryTree:
// 		init = NewBinaryTree(g, seed, init.(*BinaryTree).bias)

// 	case *Iterative:
// 		init = NewIterative(g, seed)

// 	case *Wilson:
// 		init = NewWilson(g, seed)
// 	}

// 	return init.Init()
// }
