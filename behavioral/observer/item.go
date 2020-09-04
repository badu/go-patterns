package observer

import (
	"fmt"
)

// Item represents a specific item in inventory
type Item struct {
	name      string
	quantity  int
	observers []Observer
}

// Subscribe Subscribes a customer for email notification when the item is in stock
func (i *Item) Subscribe(observer Observer) {
	fmt.Printf("Customer %s Subscribed to Notifications\n", observer.Name())
	i.observers = append(i.observers, observer)
}

// UnSubscribe Un-Subscribes a customer from email notification
func (i *Item) UnSubscribe(observer Observer) {
	for ii, o := range i.observers {
		if o.Name() == observer.Name() {
			i.observers = append(i.observers[:ii], i.observers[ii+1:]...)
			fmt.Printf("Customer %s Un-Subscribed from Notifications\n", o.Name())
		}
	}
}

// NotifyAll notifies all subscribed customers
func (i *Item) NotifyAll() {
	for _, observer := range i.observers {
		observer.SendNotification(i.name)
	}
}

// UpdateStock updates the quantity of a particular item
// Triggers a notification to all subscribed customers if the quantity>0
func (i *Item) UpdateStock(quantity int) {
	i.quantity = quantity
	if i.quantity > 0 {
		fmt.Printf("%s in stock!!! Notifying Subscribed Customers...\n", i.name)
		i.NotifyAll()
	}
}

// SetName sets the item name
func (i *Item) SetName(name string) {
	i.name = name
}
