# GoMaze: Random Maze Generation for Reinforcement Learning in Go

GoMaze provides random maze generation for reinforcement learning in `Go`.
To start, create a new `Maze`, then use the `Step()` and `Reset()` methods
to take actions and reset the maze when the agent has reached the goal
respectively. Or, if you'd like to try the mazes with learning algorithms
already implemented in `Go`, see my [GoLearn](https://github.com/samuelfneumann/GoLearn)
repository.

You can even interactively play on `Maze`s in the terminal using the
`Play()` method.

## Algorithms

The following maze generation algorithms are implemented. Each has its
own [biases](https://en.wikipedia.org/wiki/Maze_generation_algorithm),
which are important to take into account when learning on  these mazes:

1. Depth-First algorithms - biases towards long corridors
    * Backtracking recursion
    * Iterative
2. Uniformly Distributed Mazes - expensive to construct maze
    * Wilson's algorithm
    * Aldous-Broder algorithm
3. Binary Tree Algorithm - [diagonal bias](http://weblog.jamisbuck.org/2011/2/1/maze-generation-binary-tree-algorithm) with
two of the four sides of the maze being spanned by a single corridor.

## Maze Examples

### Backtracking

```go
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
    gomaze.NewBacktracking(time.Now().UnixNano()),
)
if err != nil {
    log.Fatal(err)
}

fmt.Println(m)

+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
| x     |           |       |                           |   |
+---+   +   +---+   +   +---+   +---+---+---+   +---+   +   +
|       |   |   |   |   |       |           |       |   |   |
+   +   +   +   +   +   +   +---+   +   +---+---+   +   +   +
|   |   |   |   |   |   |   |       |   |           |       |
+   +---+   +   +   +   +   +---+   +---+   +---+---+---+   +
|           |   |   |               |   |   |               |
+   +---+---+   +   +---+---+---+---+   +   +---+---+---+---+
|   |           |       |               |   |               |
+   +---+   +   +---+   +   +---+---+---+   +   +---+---+   +
|           |   |   |       |               |           |   |
+---+---+---+   +   +---+---+   +---+---+---+   +---+---+   +
|           |   |           |       |           |       |   |
+   +---+   +   +---+---+   +   +   +   +---+---+   +   +   +
|   |       |       |       |   |       |       |   |       |
+   +---+   +---+   +   +   +---+---+---+   +   +   +---+   +
|       |       |       |               |   |   |       |   |
+   +   +---+---+---+---+---+---+---+   +   +   +---+   +   +
|   |                                       |           | 🏳|
+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
```

### Binary Tree

```go
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
    gomaze.NewBinaryTree(time.Now().UnixNano()),
)
if err != nil {
    log.Fatal(err)
}

fmt.Println(m)

+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
| x                                                         |
+   +---+---+   +   +   +---+---+   +   +---+---+---+---+   +
|           |   |   |           |   |                   |   |
+   +---+---+   +---+   +   +   +---+   +---+   +---+   +   +
|           |       |   |   |       |       |       |   |   |
+   +---+---+---+---+---+---+---+---+---+   +   +   +   +   +
|                                       |   |   |   |   |   |
+   +   +   +   +---+   +---+---+   +   +   +   +   +---+---+
|   |   |   |       |           |   |   |   |   |           |
+   +   +   +---+---+---+   +   +   +---+   +---+---+---+---+
|   |   |               |   |   |       |                   |
+   +   +---+---+---+   +   +   +---+---+   +---+---+   +   +
|   |               |   |   |           |           |   |   |
+   +   +   +---+   +---+---+---+---+---+   +---+   +   +   +
|   |   |       |                       |       |   |   |   |
+   +   +   +---+---+   +   +   +---+   +---+---+---+---+   +
|   |   |           |   |   |       |                   |   |
+   +   +---+---+   +---+   +---+---+---+---+---+---+   +---+
|   |           |       |                           |     🏳|
+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
```

### Wilson's Algorithm

```go
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
)
if err != nil {
    log.Fatal(err)
}

fmt.Println(m)

+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
| x                     |       |   |                   |   |
+   +---+---+---+   +   +---+   +   +---+   +---+---+---+   +
|       |   |       |       |       |           |   |   |   |
+---+---+   +---+   +---+---+   +---+---+   +---+   +   +   +
|           |   |                   |       |       |   |   |
+   +---+   +   +   +---+   +---+   +   +---+---+   +   +   +
|       |               |   |           |   |           |   |
+   +   +---+   +   +---+---+   +---+---+   +---+   +---+   +
|   |   |   |   |   |           |               |           |
+---+---+   +---+   +   +---+   +   +---+   +---+   +---+---+
|   |               |   |   |   |   |   |                   |
+   +   +---+---+---+---+   +   +---+   +   +---+   +---+---+
|   |   |                   |   |       |   |   |           |
+   +   +---+---+---+---+   +   +   +   +---+   +   +---+   +
|   |           |                   |           |   |   |   |
+   +---+---+   +---+---+   +---+   +   +---+   +   +   +   +
|                   |           |   |   |   |       |       |
+   +---+---+   +---+   +---+   +---+   +   +---+   +---+   +
|   |               |       |   |               |       | 🏳|
+---+---+---+---+---+---+---+---+---+---+---+---+---+---+---+
```

## Acknowledgements

Inspired by [aMAZEd](https://github.com/gnmathur/aMAZEd). Some code transliterated
from this repo.
