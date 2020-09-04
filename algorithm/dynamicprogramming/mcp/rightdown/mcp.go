package rightdown

func MCPRD(grid [][]int, rLen, cLen int) int {
	if rLen == 0 {
		return 0
	}

	// initialize 2d dp array
	dp := init2DArr(rLen, cLen)

	// Cost of to reach first element of a grid = cost of
	// first element in the grid
	dp[0][0] = grid[0][0]

	// Travel in y direction
	// Cost of to reach a 0,j cell = cost cell on left + cost of current cell
	// Since only right movement is available
	for i := 1; i < cLen; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	// Travel in x direction
	// Cost of to reach a i,0 cell = cost cell on top + cost of current cell
	// Since only down movement is available
	for i := 1; i < rLen; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < rLen; i++ {
		for j := 1; j < cLen; j++ {
			// current value + min(two possible paths)
			// dp[i-1][j]: left cell
			// dp[i][j-1]: upper cell
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[rLen-1][cLen-1]
}

func min(val1, val2 int) int {
	if val1 < val2 {
		return val1
	}
	return val2
}

func init2DArr(rLen, cLen int) [][]int {
	res := make([][]int, rLen, rLen)
	for i := 0; i < rLen; i++ {
		res[i] = make([]int, cLen, cLen)
	}
	return res
}
