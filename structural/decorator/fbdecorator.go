package decorator

import "fmt"

// FBDecorator represents the facebook notifier type
type FBDecorator struct{}

// send sends a facebook notification
func (fd *FBDecorator) send(msg string) {
	fmt.Printf("Sending Facebook notification: %s\n", msg)
}
