package fibonacci

// Fib returns fibonacci series of n numbers
func Fib(n int) []int {

	r := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i < 2 {
			r = append(r, i)
			continue
		}
		r = append(r, r[i-2]+r[i-1])
	}
	return r
}
