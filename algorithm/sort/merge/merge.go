package merge

func Sort(arr []int) []int {

	aLen := len(arr)

	if aLen == 1 {
		return arr
	}

	mid := aLen / 2

	left := make([]int, mid)
	right := make([]int, aLen-mid)

	for i := 0; i < aLen; i++ {
		if i < mid {
			left[i] = arr[i]
		} else {
			right[i-mid] = arr[i]
		}
	}

	return Merge(Sort(left), Sort(right))
}

func Merge(arr1, arr2 []int) []int {

	result := make([]int, len(arr1)+len(arr2))

	i := 0
	for len(arr1) > 0 && len(arr2) > 0 {
		if arr1[0] < arr2[0] {
			result[i] = arr1[0]
			arr1 = arr1[1:]
		} else {
			result[i] = arr2[0]
			arr2 = arr2[1:]
		}
		i++
	}

	for j := 0; j < len(arr1); j++ {
		result[i] = arr1[j]
		i++
	}
	for j := 0; j < len(arr2); j++ {
		result[i] = arr2[j]
		i++
	}
	return result
}
