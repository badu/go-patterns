package maze

import "fmt"

var N = 4

// Print solution grid
func printSolution(sol [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf(" %d ", sol[i][j])
		}
		fmt.Println()
	}
}

// Check if path is available at given (x,y)
func isPathAvailable(maze [][]int, x int, y int) bool {
	// if (x, y outside maze) return false
	if x >= 0 && x < N && y >= 0 && y < N && maze[x][y] == 1 {
		return true
	}
	return false
}

// Solves the given maze using backtracking
func SolveMaze(maze [][]int) {

	sol := [][]int{{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	if !solveMaze(maze, 0, 0, sol) {
		fmt.Printf("Solution doesn't exist\n")
		return
	}
	printSolution(sol)
}

func solveMaze(maze [][]int, x int, y int, sol [][]int) bool {

	// (x,y)==destination check
	if x == N-1 && y == N-1 && maze[x][y] == 1 {
		sol[x][y] = 1
		return true
	}

	// Check if path is available
	if isPathAvailable(maze, x, y) {
		// Mark x,y in solution grid
		sol[x][y] = 1

		// Move forward in x direction
		if solveMaze(maze, x+1, y, sol) {
			return true
		}

		// Move down in y direction
		if solveMaze(maze, x, y+1, sol) {
			return true
		}

		// No further path available from x and y
		// Mark the current x, y as 0
		sol[x][y] = 0

		// Backtrack
		return false
	}
	return false
}
