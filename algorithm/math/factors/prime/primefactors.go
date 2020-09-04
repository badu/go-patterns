package prime

import (
	"math"
)

func Factors(n int) []int {
	res := make([]int, 0, 1)

	// Append the number of 2s that divide n
	for n%2 == 0 {
		res = append(res, 2)
		n = n / 2
	}

	// n must be odd at this point.
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			res = append(res, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n
	// is a prime number greater than 2 that the previous loop cannot handle
	// Ex: When n=3 after step 1
	if n > 2 {
		res = append(res, n)

		return res
	}
	return res
}
