package e1

func Pow(x, n int) int {
	res := 1
	if n == 0 {
		return 1
	}
	for i := 1; i <= n; i++ {
		res = res * x
	}
	return res
}

func PowRecur(x, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		// Check if n is even
		// 2^n = 2^(n/2) * 2^(n/2) = 2^(n/2 +n/2) = 2^(2n/2) = 2^n
		return PowRecur(x, n/2) * PowRecur(x, n/2)
	} else {
		// If x is odd multiply by x
		// 2^(n+1) = 2 * 2^(n/2) * 2^(n/2)
		return x * PowRecur(x, n/2) * PowRecur(x, n/2)
	}
}
