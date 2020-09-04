package bubble

func Sort(arr []int) {
	aLen := len(arr)
	for i := 0; i < aLen-1; i++ {
		for j := 1; j < aLen-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
