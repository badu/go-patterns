package ways

func CoinChange(coins []int, total int) int {
	rLen := len(coins) + 1
	cLen := total + 1
	dp := init2DArr(rLen, cLen)

	// for i==0 dp[i][j]=0
	for i := 1; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if j == 0 {
				//If the total amount is 0 there is one way to make that amount
				//with whatever no. of coins you have
				// ie by taking none of them
				dp[i][j] = 1
			} else if coins[i-1] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				// 1. dp[i-1][j]: Exclude ith coin
				// 2. dp[i][j-coins[i-1]]: Include ith coin
				// 3. Add 1&2
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
			}
		}
	}
	return dp[rLen-1][cLen-1]
}

func init2DArr(rLen, cLen int) [][]int {
	res := make([][]int, rLen, rLen)
	for i := 0; i < rLen; i++ {
		res[i] = make([]int, cLen, cLen)
	}
	return res
}
