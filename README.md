# GoMaze: Random Maze Generation for Reinforcement Learning in Go

GoMaze provides random maze generation for reinforcement learning in `Go`.

Currently, only the maze generation is implemented, and the overall RL
environments have not been implemented.

## Algorithms

The following maze generation algorithms are implemented. Each has its
own [biases](https://en.wikipedia.org/wiki/Maze_generation_algorithm),
which are important to take into account when learning on  these mazes:

1. Depth-First algorithms - biases towards long corridors
    * Backtracking recursion
    * Iterative
2. Uniformly Distributed Mazes
    * Wilson's algorithm
    * Aldous-Broder algorithm
3. Binary Tree Algorithm - [diagonal bias](http://weblog.jamisbuck.org/2011/2/1/maze-generation-binary-tree-algorithm) with
two of the four sides of the maze being spanned by a single corridor.

## Acknowledgements

Inspired by [aMAZEd](https://github.com/gnmathur/aMAZEd). Some code transliterated
from this repo.