package rotate

import (
	"api/algorithm/misc"
)

// Left rotates the given in slice left by n
func Left(arr []int, n int) {
	misc.Reverse(arr[:n-1])
	misc.Reverse(arr[n-1:])
	misc.Reverse(arr[:])
}
