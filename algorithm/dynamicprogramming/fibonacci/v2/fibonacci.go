package v1

func Fibo(n int) int {
	if n <= 1 {
		return n
	}

	dp := make([]int, 2, 2)

	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		tmp := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = tmp
	}
	return dp[1]
}
