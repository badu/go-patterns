package decorator

// NotifyDecorator is the decorator type
// Adds functionality(SMS, Facebook notification) to Notify type
type NotifyDecorator struct {
	Notify    *Notify
	Notifiers []Notifier
}

// send sends notification to all Notifiers(sms, facebook here)
func (nd *NotifyDecorator) send(msg string) {
	if nd.Notify != nil {
		nd.Notify.send(msg)
	}
	for _, notifier := range nd.Notifiers {
		notifier.send(msg)
	}
}

// SendNotification is the API exposed to the client
// Client calls this method to send notifications
func (nd *NotifyDecorator) SendNotification(msg string) {
	nd.send(msg)
}
