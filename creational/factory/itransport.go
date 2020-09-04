package factory

// iTransport is the transport interface
// defines all methods necessary to create any transport object
type iTransport interface {
	SetName(string)
	Name() string
}
