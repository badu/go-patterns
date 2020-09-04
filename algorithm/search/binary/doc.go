package binary

// =====================================================================================================================
// Binary Search
// =====================================================================================================================
// Binary search works on sorted arrays
// Binary Search searches a sorted array by repeatedly dividing the search interval in half.
// Begin with an interval covering the whole array.
// If the value of the search key is less than the item in the middle of the interval,
// narrow the interval to the lower half.
// Otherwise narrow it to the upper half.
// Repeatedly check until the value is found or the interval is empty.
// The idea of binary search is to use the information that the array is sorted and reduce the time complexity
//
// Time Complexity
// ===========================================
// The time complexity of binary search algorithm is O(Log n)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
// x := 23
// fmt.Printf("%d is at position %d\n", x, binary.Search(arr, x))
//
//}
// RUN INSTRUCTIONS 2:
// ===========================================
// func main() {
// arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91}
// x := 1
// fmt.Printf("%d is at position %d\n", x, binary.Recur(arr, x, 0, len(arr)-1))
//
//}
