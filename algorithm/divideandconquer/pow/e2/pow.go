package e2

func Pow(x int, n uint) int {
	if n == 0 {
		return 1
	}

	tmp := Pow(x, n/2)
	if n%2 == 0 {
		return tmp * tmp
	} else {
		return x * tmp * tmp
	}
}
