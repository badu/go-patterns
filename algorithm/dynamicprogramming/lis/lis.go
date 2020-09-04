package lis

// LIS returns the lis length for the arr array
func LIS(arr []int) int {

	aLen := len(arr)

	lis := make([]int, aLen, aLen)

	// Initialize all lis elements to 1
	// Since LIS is at least 1 for a given element
	for i := 0; i < aLen; i++ {
		lis[i] = 1
	}

	// Increase lis[i] if arr[j]<arr[i] and lis[i] < lis[j]+1
	// lis[i] < lis[j]+1 is for conditions when arr[j] < arr[i] but
	// lis[i] < lis[j]+1
	for i := 1; i < aLen; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1
			}
		}
	}

	var max int

	// Get the max element from lis array
	for i := 0; i < aLen; i++ {
		if lis[i] > max {
			max = lis[i]
		}
	}
	return max
}

// LIS returns the longest increasing subsequence for the arr array amd the lis length
func LISSequence(arr []int) (int, []int) {

	aLen := len(arr)

	res := make([][]int, aLen, aLen)
	lis := make([]int, aLen, aLen)

	// Initialize all lis elements to 1
	for i := 0; i < aLen; i++ {
		lis[i] = 1
	}

	// Increase lis[i] if arr[j]<arr[i] and lis[j]+1 > lis[i]
	for i := 1; i < aLen; i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] && lis[i] < lis[j]+1 {
				lis[i] = lis[j] + 1

				res[i] = append(res[i], arr[j])
				// Since the j loop wont be executed for i==j case
				if j == i-2 {
					res[i] = append(res[i], arr[i])
				}
			}
		}
	}

	var max, maxIndex int
	// Get the max element from lis array
	for i := 0; i < aLen; i++ {
		if lis[i] > max {
			max = lis[i]
			maxIndex = i
		}
	}
	return max, res[maxIndex]
}
