package knapsack

func Knapsack(W, P []int, cap int) int {

	rLen := len(W) + 1
	cLen := cap + 1
	dp := init2DArray(rLen, cLen)

	// j = total weight
	for i := 1; i < rLen; i++ {
		for j := 1; j < cLen; j++ {
			if W[i-1] > j {
				dp[i][j] = dp[i-1][j]

			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-W[i-1]]+P[i-1])
			}
		}
	}
	return dp[rLen-1][cLen-1]
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func init2DArray(rLen, cLen int) [][]int {
	res := make([][]int, rLen, rLen)

	for i := 0; i < rLen; i++ {
		res[i] = make([]int, cLen, cLen)
	}
	return res
}
