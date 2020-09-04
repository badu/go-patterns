package selection

func Sort(arr []int) {
	aLen := len(arr)
	for i := 0; i < aLen-1; i++ {
		minIndex := i
		for j := i + 1; j < aLen; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[minIndex], arr[i] = arr[i], arr[minIndex]
	}
}
