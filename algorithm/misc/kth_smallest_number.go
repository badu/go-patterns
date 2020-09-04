package misc

import (
	"api/algorithm/sort/selection"
)

//https://practice.geeksforgeeks.org/problems/kth-smallest-element/0
func GetKthSmallestNumber(arr []int, k int) int {
	selection.Sort(arr)

	return arr[k-1]
}
