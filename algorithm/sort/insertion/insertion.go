package insertion

func Sort(arr []int) {
	aLen := len(arr)
	for i := 0; i < aLen-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
