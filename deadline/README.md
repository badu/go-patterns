# Deadline

The deadline/timeout resiliency pattern for golang.

Creating a deadline takes one parameter: how long to wait.

```go
dl := deadline.New(1 * time.Second)

err := dl.Run(func(stopper <-chan struct{}) error {
	// do something potentially slow
	// give up when the `stopper` channel is closed (indicating a time-out)
	return nil
})

switch err {
case deadline.ErrTimedOut:
	// execution took too long, oops
default:
	// some other error
}
```
