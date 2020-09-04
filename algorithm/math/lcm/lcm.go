package lcm

import "math"

func NoGCD(a, b int) int {
	var s, l int

	l = int(math.Max(float64(a), float64(b)))
	s = int(math.Min(float64(a), float64(b)))

	for i := l; ; i += l {
		if i%s == 0 {
			return i
		}
	}
}
