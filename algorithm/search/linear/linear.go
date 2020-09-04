package linear

// Search returns the index of x using linear search
// returns -1 if not found
func Search(arr []int, x int) int {
	for i, val := range arr {
		if val == x {
			return i
		}
	}
	return -1
}
