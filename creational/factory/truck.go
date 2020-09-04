package factory

// truck represents a truck transport type
type truck struct {
	transport
}

// truck return an instance of a truck
func getTruck() iTransport {

	return &truck{}
}
