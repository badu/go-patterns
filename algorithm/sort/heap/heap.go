package heap

// Sort heap sorts the array arr
// 1. Form heap from the given array
// 2. Delete elements from the heap
// When you delete elements from the max heap
// and put the deleted elements in an array
// it will be in ascending order by default
// Elements are always deleted from the root in a heap(0th element of the array)
func Sort(arr []int) {
	aLen := len(arr)

	// In the case of a complete tree, the first index of a non-leaf node is given by n/2 - 1.
	// All other nodes after that are leaf-nodes and thus don't need to be heapified.
	for i := aLen/2 - 1; i >= 0; i-- {
		heapify(arr, aLen, i)
	}

	// Largest item is at the root (arr[0])
	// exchange it with the last index in the heap array(excluding the tails items not part of the heap)
	// Push a[0] down to its correct position by calling heapify
	for i := aLen - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		heapify(arr, i, 0)
	}

}

func heapify(arr []int, n, i int) {

	parentIndex := i
	lChildIndex := 2*i + 1
	rChildIndex := 2*i + 2

	// Compare left child with the its parent
	if lChildIndex < n && arr[lChildIndex] > arr[parentIndex] {
		parentIndex = lChildIndex
	}

	// Compare right child with its parent
	if rChildIndex < n && arr[rChildIndex] > arr[parentIndex] {
		parentIndex = rChildIndex
	}

	// From last steps parentIndex = lChildIndex, parentIndex = rChildIndex
	// So if left or right child was > parentIndex i will not be equal to parentIndex anymore
	// If not equal swap i with parentIndex index
	// Recursively continue until the structure(array) becomes a heap
	if parentIndex != i {
		arr[i], arr[parentIndex] = arr[parentIndex], arr[i]

		heapify(arr, n, parentIndex)
	}
}
