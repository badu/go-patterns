package merge

// =====================================================================================================================
// Merge Sort
// =====================================================================================================================
// Like QuickSort, Merge Sort is a Divide and Conquer algorithm.
// It divides input array in two halves, calls itself for the two halves and then merges the two sorted halves
// The merge() function is used for merging two halves.
//
// Merge sort utilizes divide-and-conquer rule to break the problem into sub-problems,
// the problem in this case being, sorting a given array.
// In merge sort, we break the given array midway, for example if the original array had 6 elements,
// then merge sort will break it down into two subarrays with 3 elements each.
// But breaking the original array into 2 smaller subarrays is not helping us in sorting the array.
// So we will break these subarrays into even smaller subarrays,
// until we have multiple subarrays with single element in them.
// Now, the idea here is that an array with a single element is already sorted,
// so once we break the original array into subarrays which has only a single element,
// we have successfully broken down our problem into base problems.
// And then we have to merge all these sorted subarrays, step by step to form one single sorted array.
// Let's consider an array with values {14, 7, 3, 12, 9, 11, 6, 12}
// Refer image1, we have a pictorial representation of how merge sort will sort the given array.
//
// Applications
// ===========================================
// 1. Merge Sort is useful for sorting linked lists in O(n Log n) time.In the case of linked lists,
//    the case is different mainly due to the difference in memory allocation of arrays and linked lists.
//    Unlike arrays, linked list nodes may not be adjacent in memory. Unlike an array, in the linked list,
//    we can insert items in the middle in O(1) extra space and O(1) time.
//    Therefore merge operation of merge sort can be implemented without extra space for linked lists.
//
//    In arrays, we can do random access as elements are contiguous in memory.
//    Let us say we have an integer (4-byte) array A and let the address of A[0] be x then to access A[i],
//    we can directly access the memory at (x + i*4). Unlike arrays, we can not do random access in the linked list.
//    Quick Sort requires a lot of this kind of access.
//    In linked list to access i’th index, we have to travel each and
//    every node from the head to i’th node as we don’t have a continuous block of memory.
//    Therefore, the overhead increases for quicksort.
//    Merge sort accesses data sequentially and the need of random access is low.
//
// 2. Inversion Count Problem (https://www.geeksforgeeks.org/counting-inversions/)
// 3. Used in External Sorting (http://en.wikipedia.org/wiki/External_sorting)
//
// Time Complexity
// ===========================================
// Worst Case Time Complexity [ Big-O ]: O(n*log n)
// Best Case Time Complexity [Big-omega]: O(n*log n)
// Average Time Complexity [Big-theta]: O(n*log n)
// Auxiliary Space: O(n)
// Stable: Yes
// Time complexity is O(n Log n)
// (Got by solving the equation T(n/2) + T(n/2) + n)
// + n, because you need to merge n/2 and n/2
//
// As we have already learned in Binary Search that whenever we divide a number into half in every step,
// it can be represented using a logarithmic function,
// which is log n and the number of steps can be represented by log n + 1(at most)
// Also, we perform a single step operation to find out the middle of any subarray, i.e. O(1).
// And to merge the subarrays,
// made by dividing the original array of n elements, a running time of O(n) will be required.
// Hence the total time for mergeSort function will become n(log n + 1), which gives us a time complexity of O(n*log n).
// It is the best Sorting technique used for sorting Linked Lists.
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{95, 80, 30, 90, 50, 40, 70, 10}
// res := merge.Sort(arr)
// fmt.Println("Result: ", res)
//}
// Output
// ===========================================
// Result:  [10 30 40 50 70 80 90 95]
