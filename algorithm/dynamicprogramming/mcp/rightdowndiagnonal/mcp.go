package rightdowndiagnonal

func MCPRDD(grid [][]int, rLen, cLen int) int {

	if cLen == 0 {
		return 0
	}

	dp := init2DArr(rLen, cLen)

	dp[0][0] = grid[0][0]

	for i := 1; i < cLen; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}

	for i := 1; i < rLen; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}

	for i := 1; i < rLen; i++ {
		for j := 1; j < cLen; j++ {
			dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
		}
	}
	return dp[rLen-1][cLen-1]
}

func min(val1, val2, val3 int) int {
	if val1 < val2 && val1 < val3 {
		return val1
	} else if val2 < val1 && val2 < val3 {
		return val2
	}
	return val3
}

func init2DArr(rLen, cLen int) [][]int {
	res := make([][]int, rLen, rLen)
	for i := 0; i < rLen; i++ {
		res[i] = make([]int, cLen, cLen)
	}
	return res
}
