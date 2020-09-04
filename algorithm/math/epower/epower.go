package epower

func Pow(n int, x float64) float64 {

	sum := 1.0
	for i := n - 1; i > 0; i-- {
		sum = 1 + x*sum/float64(i)
	}
	return sum
}
