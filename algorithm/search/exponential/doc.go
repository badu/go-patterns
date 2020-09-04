package exponential

// =====================================================================================================================
// Exponential Search
// =====================================================================================================================
// The name of this searching algorithm may be misleading as it works in O(Log n) time
// The name comes from the way it searches an element.
// It searches for an element in a sorted array by jumping 2^i elements every iteration where i represents
// the value of loop control variable, and then verifying if
// the search element is present between last jump and the current jump
// Exponential search involves two steps:
// 1. Find range where element is present
// 2. Do Binary Search in above found range.
// Exponential search is also known as doubling or galloping search or finger search.
// If L and U are the upper and lower bound of the list, then L and U both are the power of 2.
// For the last section, the U is the last position of the list. For that reason, it is known as exponential.
// Often confused because of the name, the algorithm is named so not because of the time complexity.
// The name arises as a result of the algorithm jumping elements with steps equal to exponents of 2
//
// How to find the range where element may be present?
// ===========================================
// The idea is to start with subarray size 1, compare its last element with x,
// then try size 2, then 4 and so on until last element of a subarray is not greater.
// Once we find an index i (after repeated doubling of i),
// we know that the element must be present between i/2 and i
// (Why i/2? because we could not find a greater value in previous iteration)
//
// Time Complexity
// ===========================================
// The time(worst time) complexity of exponential search algorithm is O(Log i)
// Average case time complexity: O(log i)
// Best case time complexity: O(1)
// Auxiliary space is O(1)
//
// When to use
// ===========================================
// 1. Exponential Search is particularly useful for unbounded searches, where size of array is infinite
//
// Pros
// ===========================================
// 1. It works better than Binary Search for bounded arrays,
//    when the element to be searched is closer to the first element.
//    This is because exponential search runs in O(Log i) time, where i is the index of the number being searched
//    for in the array, whereas binary search would run in O(Log n) time where n is the number of elements in the array
//
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91, 95, 96, 97, 98, 99, 100, 103, 105, 107}
// x := 2
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 5
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 8
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 12
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 16
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 23
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 38
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 56
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 72
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 91
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 95
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 96
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 97
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 98
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 99
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 100
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
//
// x = 22
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
//
// x = 3
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 102
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 103
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 105
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 107
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
// x = 108
// fmt.Printf("%d is at position %d\n", x, exponential.Search(arr, x))
//
//}
