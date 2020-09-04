package selection

// =====================================================================================================================
// Selection Sort
// =====================================================================================================================
// The selection sort algorithm sorts an array by repeatedly finding the minimum element (considering ascending order)
// from unsorted part and putting it at the beginning.
// The algorithm maintains two sub-arrays in a given array.
// 1) The subarray which is already sorted.
// 2) Remaining subarray which is unsorted.
// In every iteration of selection sort,
// the minimum element (considering ascending order) from the unsorted subarray is picked and
// moved to the sorted subarray.
// This algorithm will first find the smallest element in the array and swap it with the element in the first position,
// then it will find the second smallest element and swap it with the element in the second position,
// and it will keep on doing this until the entire array is sorted.
// It is called selection sort because it repeatedly selects the next-smallest element and
// swaps it into the right place.
//
// Example
// ===========================================
// arr[] = 64 25 12 22 11
// Find the minimum element in arr[0...4] and place it at beginning
// 11 25 12 22 64
//
// Find the minimum element in arr[1...4] and place it at beginning of arr[1...4]
// 11 12 25 22 64
//
// Find the minimum element in arr[2...4] and place it at beginning of arr[2...4]
// 11 12 22 25 64
//
// Find the minimum element in arr[3...4] and place it at beginning of arr[3...4]
// 11 12 22 25 64
//
// Time Complexity
// ===========================================
// Worst Case Time Complexity [ Big-O ]: O(n^2)
// Best Case Time Complexity [Big-omega]: O(n^2)
// Average Time Complexity [Big-theta]: O(n^2)
// Auxiliary Space: O(1)
// The good thing about selection sort is it never makes more than O(n) swaps and
// can be useful when memory write is a costly operation.
// Stability: The default implementation is not stable. However it can be made stable.
// In Place : Yes, it does not require extra space.
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// arr := []int{95, 80, 30, 90, 50, 40, 70, 10}
// selection.Sort(arr)
// fmt.Println("Result: ", arr)
//
//}
//
// Output
// ===========================================
// Result:  [10 30 40 50 70 80 90 95]
