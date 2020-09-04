package counting

func Sort(arr []int, k int) []int {

	aLen := len(arr)

	count := make([]int, k+1, k+1)
	for i := 0; i < aLen; i++ {
		count[arr[i]]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	res := make([]int, aLen, aLen)
	// Begin from the right  of count[] to maintain stability of the algorithm
	for i := aLen - 1; i >= 0; i-- {
		count[arr[i]]--
		res[count[arr[i]]] = arr[i]
	}
	return res
}

func SortNegative(arr []int, min, max int) []int {

	aLen := len(arr)

	count := make([]int, max-min+1, max-min+1)
	for i := 0; i < aLen; i++ {
		count[arr[i]-min]++
	}

	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	res := make([]int, aLen, aLen)
	for i := aLen - 1; i >= 0; i-- {
		count[arr[i]-min]--
		res[count[arr[i]-min]] = arr[i]
	}
	return res
}
