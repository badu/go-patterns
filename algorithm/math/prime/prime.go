package prime

func Eratosthenes(n int) []int {
	res := make([]int, 0, 1)

	pMap := make(map[int]bool)
	for i := 2; i < n; i++ {
		pMap[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if pMap[p] {
			for j := p * p; j <= n; j += p {
				pMap[j] = false
			}
		}
	}

	for i := 2; i <= n; i++ {
		if pMap[i] {
			res = append(res, i)
		}
	}
	return res
}
