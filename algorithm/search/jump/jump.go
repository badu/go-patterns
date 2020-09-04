package jump

import (
	"api/algorithm/search/linear"
	"math"
)

func Search(arr []int, x int) int {

	aLen := len(arr)
	step := int(math.Sqrt(float64(aLen)))

	return jumpSearch(arr, x, 0, step-1, step)
}

func jumpSearch(arr []int, x, l, r, step int) int {
	if r >= l {
		if x == arr[r] {
			return r
		} else if x < arr[r] {
			lPos := linear.Search(arr[l:r+1], x)
			if lPos >= 0 {
				return lPos + l
			}
			return -1
		} else if r+step >= len(arr) {
			lPos := linear.Search(arr[r+1:], x)
			if lPos >= 0 {
				return lPos + r + 1
			}
			return -1
		} else {
			return jumpSearch(arr, x, r+1, r+step, step)
		}
	}
	return -1
}
