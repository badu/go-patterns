package subsetsum

import "fmt"

func SubsetSum(arr []int, sum int) bool {

	rLen := len(arr) + 1
	cLen := sum + 1
	dp := init2DArr(rLen, cLen)

	for i := 0; i < rLen; i++ {
		dp[i][0] = true
	}

	for i := 1; i < rLen; i++ {
		for j := 1; j < cLen; j++ {
			if j < arr[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i-1]]
			}
		}
	}
	print(dp, arr, rLen, cLen)
	return dp[rLen-1][cLen-1]
}

func print(dp [][]bool, arr []int, rLen, cLen int) {
	res := make([]int, 0, 1)
	j := cLen - 1
	for i := rLen - 1; i > 0; i-- {
		if dp[i][j] != dp[i-1][j] {
			res = append(res, arr[i-1])
			j = j - arr[i-1]
		}
	}
	fmt.Println("Subset: ", res)
}

func init2DArr(rLen, cLen int) [][]bool {
	res := make([][]bool, rLen, rLen)

	for i := 0; i < rLen; i++ {
		res[i] = make([]bool, cLen, cLen)
	}
	return res
}
