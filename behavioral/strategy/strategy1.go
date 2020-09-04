package strategy

import "fmt"

// BusRoute represents a bus route type
type BusRoute struct{}

// navigate shows a bus route from start to end
func (b *BusRoute) navigate(start, end string) {
	fmt.Printf("Showing bus route from %s to %s\n", start, end)
}
