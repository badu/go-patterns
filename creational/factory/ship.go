package factory

// ship represents a ship transport type
type ship struct {
	transport
}

// ship return an instance of a ship
func getShip() iTransport {

	return &ship{}
}
