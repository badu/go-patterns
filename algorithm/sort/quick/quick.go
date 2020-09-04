package quick

// Sort Quick Sorts the given array
func Sort(arr []int, low, high int) {
	if low < high {

		// pi is partitioning index, arr[pi] is now at right place
		pi := partition(arr, low, high)

		// Recursively sort elements before pi and after pi
		Sort(arr, low, pi-1)  // Before pi
		Sort(arr, pi+1, high) // After pi
	}
}

// partition takes last element as pivot,
// places the pivot element at its correct
// position in sorted array, and places all
// smaller (smaller than pivot) to left of
// pivot and all greater elements to right
// of pivot
func partition(arr []int, low, high int) int {

	pivot := arr[high]

	// index of smaller element
	i := low - 1

	for j := low; j < high; j++ {
		// If current element is smaller than the pivot
		if arr[j] < pivot {
			i++
			// swap arr[i] and arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	// swap arr[i+1] and arr[high] (or pivot)
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
