package decorator

import "fmt"

// Notify represents the email notifier type
type Notify struct{}

// send sends an email notification
func (n *Notify) send(msg string) {
	fmt.Printf("Sending email notification: %s\n", msg)
}

// SendNotification is the API exposed to the client
// Client calls this method to send notifications
func (n *Notify) SendNotification(msg string) {
	n.send(msg)
}
