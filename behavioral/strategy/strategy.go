package strategy

// strategy is an interface for route strategy
type strategy interface {
	navigate(start, end string)
}
