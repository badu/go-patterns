package overlap

type Point struct {
	X, Y int
}

// Returns true if two rectangles (l1, r1) and (l2, r2) overlap
func DoOverlap(l1, r1, l2, r2 Point) bool {
	// If one rectangle is on left side of other
	if l1.X >= r2.X || l2.X >= r1.X {
		return false
	}

	// If one rectangle is above other
	if l1.Y <= r2.Y || l2.Y <= r1.Y {
		return false
	}
	return true
}
