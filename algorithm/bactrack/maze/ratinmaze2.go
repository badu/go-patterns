package maze

import "fmt"

var N2 = 4

// Print solution grid
func printSolution2(sol [][]int) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf(" %d ", sol[i][j])
		}
		fmt.Println()
	}
}

// Solves the given maze using backtracking
func SolveMazeRDD(maze [][]int) {

	sol := [][]int{{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0}}
	if !solveMazeRDD(maze, 0, 0, sol) {
		fmt.Printf("Solution doesn't exist\n")
		return
	}
	printSolution(sol)
}

func solveMazeRDD(maze [][]int, x int, y int, sol [][]int) bool {

	// (x,y)==destination check
	if x == N-1 && y == N-1 && maze[x][y] == 1 {
		sol[x][y] = 1
		return true
	}

	// Check if path is available
	if isPathAvailable(maze, x, y) {
		// Mark x,y in solution grid
		sol[x][y] = 1

		// Move down
		if solveMazeRDD(maze, x+1, y, sol) {
			return true
		}

		// Move right
		if solveMazeRDD(maze, x, y+1, sol) {
			return true
		}

		// Move diagonal
		if solveMazeRDD(maze, x+1, y+1, sol) {
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
