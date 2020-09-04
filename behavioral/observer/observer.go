package observer

// Observer interface
// All observers need to implement this interface
// Here the Customer type implements Observer
type Observer interface {
	SendNotification(item string)
	Name() string
}
