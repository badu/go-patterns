package editdistance

// Converts the y axis string to x axis string
func MinimumEdits(src, trgt string) int {

	rLen := len(src) + 1
	cLen := len(trgt) + 1
	dp := init2DArr(rLen, cLen)

	for i := 0; i < rLen; i++ {
		for j := 0; j < cLen; j++ {
			if i == 0 {
				dp[i][j] = j
				continue
			} else if j == 0 {
				dp[i][j] = i
				continue
			} else if src[i-1] == trgt[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
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
