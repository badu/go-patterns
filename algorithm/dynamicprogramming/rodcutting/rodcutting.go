package rodcutting

func MaxProfit(totLen int, lens, profits []int) int {
	rLen := len(lens) + 1
	cLen := totLen + 1

	dp := init2DArr(rLen, cLen)

	for i := 1; i < rLen; i++ {
		for j := 1; j < cLen; j++ {
			if lens[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-lens[i-1]]+profits[i-1])
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

func init2DArr(rLen, cLen int) [][]int {
	res := make([][]int, rLen, rLen)
	for i := 0; i < rLen; i++ {
		res[i] = make([]int, cLen, cLen)
	}
	return res
}
