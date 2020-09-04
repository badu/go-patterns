package radix

func Sort(arr []int) {
	// Find the maximum number to know number of digits
	m := max(arr)

	// Starting for lsb check every exponent
	// if m=802
	// for exponent=1 m/exponent =802
	// for exponent=10 m/exponent =80
	// for exponent=100 m/exponent =8
	for exponent := 1; m/exponent > 0; exponent = exponent * 10 {
		CountSort(arr, m, exponent)
	}
}

func CountSort(arr []int, k, exponent int) {

	aLen := len(arr)

	count := make([]int, k+1, k+1)
	for i := 0; i < aLen; i++ {
		count[(arr[i]/exponent)%10]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	res := make([]int, aLen, aLen)
	for i := aLen - 1; i >= 0; i-- {
		count[(arr[i]/exponent)%10]--
		res[count[(arr[i]/exponent)%10]] = arr[i]
	}

	for i := 0; i < aLen; i++ {
		arr[i] = res[i]
	}
}

func max(arr []int) int {
	var max int
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}
