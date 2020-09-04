package strategy

import "fmt"

// WalkRoute represents a walk route type
type WalkRoute struct{}

// navigate shows a walk route from start to end
func (w *WalkRoute) navigate(start, end string) {
	fmt.Printf("Showing walk route from %s to %s\n", start, end)
}
