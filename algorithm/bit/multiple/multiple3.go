package multiple

import "math"

func Multiple3(n int) int {

	oddCount := 0
	evenCount := 0

	// Make no. positive if +n is multiple of 3 then is -n.
	// We are doing this to avoid stack overflow in recursion
	if n < 0 {
		// Makes n positive
		n = -n
	}

	// Multiples of 3 eventually result in n=0
	if n == 0 {
		return 1
	}

	// Non multiples of 3 eventually end up in n=1
	if n == 1 {
		return 0
	}

	for n != 0 {
		// If odd bit is set then increment odd counter
		// Checking last(right most) bit
		if (n & 1) != 0 {
			oddCount++
		}

		// If even bit is set then increment even counter
		// Checking second last bit
		if (n & 2) != 0 {
			evenCount++
		}
		// Right shift n by 2
		// Since last 2 bits are already checked in the previous steps
		n = n >> 2
	}
	// Check if difference is also a multiple of 3
	// Recurse until n==1(not a multiple of 3) or n==0(multiple of 3)
	return Multiple3(int(math.Abs(float64(oddCount - evenCount))))
}
