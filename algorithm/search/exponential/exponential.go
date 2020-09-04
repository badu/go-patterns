package exponential

import (
	"api/algorithm/search/binary"
)

func Search(arr []int, x int) int {

	if len(arr) == 0 {
		return -1
	}
	return exponentialSearch(arr, x, 0, 0)
}

func exponentialSearch(arr []int, x, l, r int) int {
	if r >= l {
		if x == arr[r] {
			return r
		} else if x < arr[r] {
			bPos := binary.Search(arr[l:r+1], x)
			if bPos >= 0 {
				return bPos + l
			}
			return -1
		} else if r*2 >= len(arr) {
			bPos := binary.Search(arr[r+1:], x)
			if bPos >= 0 {
				return bPos + r + 1
			}
			return -1
		} else {
			return exponentialSearch(arr, x, r+1, (r+1)*2)
		}
	}
	return -1
}
