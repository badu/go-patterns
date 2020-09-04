package base2

func Log2n(n int) int {
	if n > 1 {
		return 1 + Log2n(n/2)
	}
	return 0
}
