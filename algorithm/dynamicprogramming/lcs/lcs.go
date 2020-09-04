package lcs

func LCS(s1, s2 string) int {

	rLen := len(s1) + 1
	cLen := len(s2) + 1

	lcs := init2DArr(rLen, cLen)

	for i := 1; i < rLen; i++ {
		for j := 1; j < cLen; j++ {
			if s1[i-1] == s2[j-1] {
				lcs[i][j] = 1 + lcs[i-1][j-1]
			} else {
				lcs[i][j] = max(lcs[i][j-1], lcs[i-1][j])
			}
		}
	}
	return lcs[rLen-1][cLen-1]
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

func init2DArr(rowLen, colLen int) [][]int {

	dp := make([][]int, rowLen, rowLen)

	for i := 0; i < rowLen; i++ {
		dp[i] = make([]int, colLen, colLen)
	}
	return dp
}
