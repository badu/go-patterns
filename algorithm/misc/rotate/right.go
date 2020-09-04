package rotate

import (
	"api/algorithm/misc"
)

// Right rotates the given int slice right by n
func Right(arr []int, n int) {
	misc.Reverse(arr[n:])
	misc.Reverse(arr[:n])
	misc.Reverse(arr[:])
}
