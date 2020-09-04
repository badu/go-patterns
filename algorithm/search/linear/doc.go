package linear

// =====================================================================================================================
// Linear Search
// =====================================================================================================================
// Linear search is one of the simplest search algorithms
// For a given array arr linear search does the following steps:
// 1. Start from the leftmost element of arr[] and one by one compare x with each element of arr[]
// 2. If x matches with an element, return the index.
// 3. If x doesn’t match with any of elements, return -1.
// Linear search is rarely used practically because other search algorithms such as
// the binary search algorithm and hash tables allow significantly faster searching compared to Linear search
//
// Time Complexity
// ===========================================
// The time complexity of linear search algorithm is O(n)
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// arr := []int{1, 10, 30, 15}
// x := 30
// fmt.Printf("%d is at position %d\n", x, linear.Search(arr, x))
//
//}
