package strategy

import "fmt"

// CarRoute represents a car route type
type CarRoute struct{}

// navigate shows you a car route from start to end
func (c *CarRoute) navigate(start, end string) {
	fmt.Printf("Showing car route from %s to %s\n", start, end)
}
