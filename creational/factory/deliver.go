package factory

import "fmt"

// Deliver delivers the item by the specified transport mode
func Deliver(item string, transport iTransport) {
	fmt.Printf("Delivering %s by %s\n", item, transport.Name())
}
