# retrier

Creating a retrier takes two parameters:
- the times to back-off between retries (and implicitly the number of times to
  retry)
- the classifier that determines which errors to retry

```go
r := retrier.New(retrier.ConstantBackoff(3, 100*time.Millisecond), nil)

err := r.Run(func() error {
	// do some work
	return nil
})

if err != nil {
	// handle the case where the work failed three times
}
```
