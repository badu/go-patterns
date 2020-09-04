package strategy

// Navigator represents the navigator type
// Has methods to set and navigate through different strategies
type Navigator struct {
	strategy strategy
}

// SetStrategy sets the route strategy
func (n *Navigator) SetStrategy(strategy strategy) {
	n.strategy = strategy
}

// Navigate will show you a route to navigate from start to end based on the navigate strategy
func (n *Navigator) Navigate(start, end string) {
	n.strategy.navigate(start, end)
}
