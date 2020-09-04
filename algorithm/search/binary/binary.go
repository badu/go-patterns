package binary

// Search returns the index of x using binary search
// returns -1 if not found
func Search(arr []int, x int) int {
	l, r := 0, len(arr)-1
	for r >= l {
		mid := (l + r) / 2
		if x == arr[mid] {
			return mid
		} else if x < arr[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// Search returns the index of x using binary search
// returns -1 if not found
// Uses recursion
func Recur(arr []int, x, l, r int) int {
	mid := (l + r) / 2
	if r >= l {
		if x == arr[mid] {
			return mid
		} else if x < arr[mid] {
			return Recur(arr, x, l, mid-1)
		} else {
			return Recur(arr, x, mid+1, r)
		}
	}
	return -1
}
