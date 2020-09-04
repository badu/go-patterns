package decorator

// Notifier is the common interface for notifications
type Notifier interface {
	send(msg string)
}
