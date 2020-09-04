package counting

// =====================================================================================================================
// Counting Sort
// =====================================================================================================================
// https://youtu.be/pEJiGC-ObQE
// Counting Sort Algorithm is an efficient sorting algorithm that
// can be used for sorting elements within a specific range.
// Counting sort is a sorting technique based on keys between a specific range.
// Counting sort is a sorting algorithm that sorts the elements of an array by counting the
// number of occurrences of each unique element in the array.
// The count is stored in an auxiliary array and the sorting is done by mapping the count as an index of
// the auxiliary array.
// It works by counting the number of objects having distinct key values (kind of hashing).
// Then doing some arithmetic to calculate the position of each object in the output sequence.
// Counting sort is a stable sort
// Counting sort for ascii characters: https://www.geeksforgeeks.org/counting-sort/
//
// Counting sort can also be extended for negative numbers:
// One way you can handle such cases is by doing a manual mapping.
// Say you have array elements as : -9, 1, 2, 7, 3, 6, -1
// Most negative element is -9.
// So take index for -9 as 0(ie. -9+ 9)
// Take index for 1 as 1+9=10
// index for 2 as 2+9=11 and so on.
// So by doing such index manipulation you can handle the negative numbers.
// Counting Sort algorithm is efficient if the range of input data (k) is
// not much greater than the number of elements in the input array (n).
// It is an integer-based sorting algorithm unlike others which are usually comparison-based.
// A comparison-based sorting algorithm sorts numbers only by comparing pairs of numbers.
// Few examples of comparison based sorting algorithms are quick sort, merge sort, bubble sort, selection sort,
// heap sort, insertion sort, whereas algorithms like radix sort, bucket sort and
// comparison sort fall into the category of non-comparison based sorting algorithms.
//
// Key Points
// ===========================================
// 1. Counting sort is efficient if the range of input data is not significantly greater than
//    the number of objects to be sorted.
//    Consider the situation where the input sequence is between range 1 to 10K and the data is 10, 5, 10K, 5K.
// 2. It is not a comparison based sorting.
//    Its running time complexity is O(n) with space proportional to the range of data.
// 3. It is often used as a sub-routine to another sorting algorithm like radix sort.
// 4. Counting sort uses a partial hashing to count the occurrence of the data object in O(1).
// 5. Counting sort can be extended to work for negative inputs also.
//
// Example
// ===========================================
// For simplicity, consider the data in the range 0 to 9.
// Input data: 1, 4, 1, 2, 7, 5, 2
//
// 1) Take a count array to store the count of each unique object.
// Index:     0  1  2  3  4  5  6  7  8  9
// Count:     0  2  2  0  1  1  0  1  0  0
//
// 2) Modify the count array such that each element at each index stores the sum of previous counts.
// Index:     0  1  2  3  4  5  6  7  8  9
// Count:     0  2  4  4  5  6  6  7  7  7
//
// 3) The modified count array indicates the position of each object in the output sequence.
//    Output each object from the input sequence followed by decreasing its count by 1.
// Process the input data: 1, 4, 1, 2, 7, 5, 2.
// Do it from right to left to maintain the stability of the algorithm
// Position of 2 is count[2]-1= 3
// Put data 2 at index 3 in output.
// -1 because count array has count values, but we need index for output array
//
// Time Complexity
// ===========================================
// Time complexity is O(n+k) where n is the number of elements in input array and k is the range of input.
// Worst Case Complexity: O(n+k)
// Best Case Complexity: O(n+k)
// Average Case Complexity: O(n+k)
// Auxiliary Space: O(n+k)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{-10, -5, 1, 4, 1, 2, 7, 5, 2}
// res := counting.SortNegative(arr, -10, 7)
//}
// Output
// ===========================================
// [-10 -5 1 1 2 2 4 5 7]
//
// RUN INSTRUCTIONS 2:
// ===========================================
// func main() {
// arr := []int{1, 4, 1, 2, 7, 5, 2}
// res := counting.Sort(arr, 7)
//}
// Output
// ===========================================
// [1 1 2 2 4 5 7]
