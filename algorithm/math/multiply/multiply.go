package multiply

func NoOperator(x, y int) int {

	if y == 0 {
		return 0
	}

	if y > 0 {
		return x + NoOperator(x, y-1)
	}

	if y < 0 {
		return -NoOperator(x, -y)
	}
	return -1
}
