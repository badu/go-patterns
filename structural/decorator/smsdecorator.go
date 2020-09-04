package decorator

import "fmt"

// SmsDecorator represents the SMS notifier type
type SmsDecorator struct{}

// send sends SMS notification
func (sd *SmsDecorator) send(msg string) {
	fmt.Printf("Sending SMS notification: %s\n", msg)
}
