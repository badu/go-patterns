package lcm

func ByGCD(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Euclid’s algorithm for GCD
func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}
