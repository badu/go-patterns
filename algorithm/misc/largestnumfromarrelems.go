package misc

import (
	"sort"
	"strconv"
)

// https://practice.geeksforgeeks.org/problems/largest-number-formed-from-an-array/0
// LargestNumArrElems returns the largest number possible from the given array elements
func LargestNumArrElems(arr []int) string {

	sort.Ints(arr)

	var res string
	for i := len(arr) - 1; i >= 0; i-- {
		res += strconv.Itoa(arr[i])
	}
	return res
}
