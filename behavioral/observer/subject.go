package observer

// Subject represents the observable type
// All type that needs to be observed must implement this interface
// Here the Item type implements this interface
type Subject interface {
	Subscribe(observer Observer)
	UnSubscribe(observer Observer)
	NotifyAll()
}
