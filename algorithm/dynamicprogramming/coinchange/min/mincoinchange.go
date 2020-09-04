package min

func CoinChange(coins []int, total int) int {
	rLen := len(coins)
	cLen := total + 1
	dp := init2DArr(rLen, cLen)

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if j == 0 {
				//If the total is 0, the min no. of coins to make that total is 0
				dp[i][j] = 0
			} else if i == 0 {
				dp[i][j] = j
			} else if coins[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 1. dp[i-1][j]: Exclude ith coin
				// 2. 1+dp[i][j-coins[i]]: Include the ith coin
				//    1--> indicates 1 coin included (current coin)+
				//    after including i total(remaining amount) becomes j-ith coin
				// Take the minimum of 1&2
				dp[i][j] = min(dp[i-1][j], 1+dp[i][j-coins[i]])
			}
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
