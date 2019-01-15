package main

func main() {
	// creation end channel for determine of finishing of calculations
	endChan := make(chan interface{})

	// start processing
	start(endChan)
	<-endChan
}
