package lps

// LPS returns the length of the longest palindromic substring of str
// Refer doc.go for explanation
func LPS(str string) int {
	n := len(str)
	var i, j, cl int

	// Initialize DP array
	dp := init2DArray(n, n)

	// Strings of length 1 are palindrome of length 1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}

	// Build the table. Note that the lower diagonal values of table are
	// useless and not filled in the process. The values are filled in a
	// manner similar to Matrix Chain Multiplication DP solution (See
	// https://www.geeksforgeeks.org/matrix-chain-multiplication-dp-8/). cl is length of substring
	for cl = 2; cl <= n; cl++ {
		for i = 0; i < n-cl+1; i++ {
			j = i + cl - 1
			if str[i] == str[j] && cl == 2 {
				dp[i][j] = 2
			} else if str[i] == str[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}

		}
	}

	return dp[0][n-1]
}

func init2DArray(rLen, cLen int) [][]int {
	res := make([][]int, rLen, rLen)

	for i := 0; i < rLen; i++ {
		res[i] = make([]int, cLen, cLen)
	}
	return res
}

func max(val1, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}

// RUN INSTRUCTIONS (LPS):
// ===========================================
// str := "GEEKS FOR GEEKS"
// fmt.Printf("LPS of %s: %d\n", str, lps.LPS(str))
//
// str = "BBABCBCAB"
// fmt.Printf("LPS of %s: %d\n", str, lps.LPS(str))
//
// Output
// ===========================================
// LPS of GEEKS FOR GEEKS: 7
// LPS of GEEKS FOR GEEKS: 7
