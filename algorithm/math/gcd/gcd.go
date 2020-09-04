package gcd

func GCD(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else if a == b {
		return a
	}

	if a > b {
		return GCD(a-b, b)
	}
	return GCD(a, b-a)
}
