package factory

// transport represents the transport type
// Implements the iTransport type
type transport struct {
	name string
}

// SetType sets the transportation name
func (t *transport) SetName(name string) {
	t.name = name
}

// Type returns the transportation name
func (t *transport) Name() string {
	return t.name
}
