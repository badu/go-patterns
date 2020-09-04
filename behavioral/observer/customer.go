package observer

import "fmt"

// Customer represents the customer type
type Customer struct {
	name string
}

// SendNotification sends an email to the customer
func (c *Customer) SendNotification(item string) {
	fmt.Printf("%s is in stock. Sending email to %s\n", item, c.Name())
}

// Name returns the name of the customer
func (c *Customer) Name() string {
	return c.name
}

// SetName sets the customer name
func (c *Customer) SetName(name string) {
	c.name = name
}
