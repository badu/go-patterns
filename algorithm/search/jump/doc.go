package jump

// =====================================================================================================================
// Jump Search
// =====================================================================================================================
// Like Binary Search, Jump Search is a searching algorithm for sorted arrays.
// The basic idea is to check fewer elements (than linear search) by jumping ahead by
// fixed steps or skipping some elements in place of searching all elements.
// For example, suppose we have an array arr[] of size n and block (to be jumped) size m.
// Then we search at the indexes arr[0], arr[m], arr[2m]…..arr[km] and so on.
// Once we find the interval (arr[km] < x < arr[(k+1)m]),
// we perform a linear search operation from the index km to find the element x.
// Let’s consider the following array: (0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610).
// Length of the array is 16.
// Jump search will find the value of 55 with the following steps assuming that the block size to be jumped is 4.
// STEP 1: Jump from index 0 to index 4;
// STEP 2: Jump from index 4 to index 8;
// STEP 3: Jump from index 8 to index 12;
// STEP 4: Since the element at index 12 is greater than 55 we will jump back a step to come to index 8.
// STEP 5: Perform linear search from index 8 to get the element 55.
// It is also called block search algorithm
// Binary Search is better than Jump Search,
// but Jump search has an advantage that we traverse back only once (Binary Search may require up to O(Log n) jumps,
// consider a situation where the element to be searched is the smallest element or smaller than the smallest).
// So in a system where binary search is costly, we use Jump Search
//
// What is the optimal block size to be skipped?
// ===========================================
// If n is the no. of elements and m is the block size, in the worst case,
// we have to do n/m jumps, if there are x left over elements at the end after the jumps(x<m),
// we perform m-1(in the worst case) comparisons more for linear search.
// Therefore the total number of comparisons in the worst case will be ((n/m) + m-1).
// The value of the function ((n/m) + m-1) will be minimum when m = √n.
// Therefore, the best step size is m = √n.
//
// To find the minimum value you need to take the differentiation w.r.t m, and by equating it to 0 we get:
// n/-(m^2)+1 = 0
//
// n/m^2 = 1
//
// m = √n
//
// Time Complexity
// ===========================================
// The time complexity of jump search algorithm is O(√n) since the max no. of jumps will be √n if the element you are
// trying to find is not in the given array.
// Auxiliary space is O(1)
//
// Pros
// ===========================================
// 1. It is faster than the linear search technique which has a time complexity of O(n) for searching an element
//
// Cons
// ===========================================
// 1. It is slower than binary search algorithm which searches an element in O(log n)
// 2. It requires the input array to be sorted
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91, 95, 96, 97, 98, 99, 100, 103, 105, 107}
// x := 2
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 5
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 8
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 12
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 16
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 23
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 38
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 56
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 72
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 91
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 95
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 96
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 97
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 98
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 99
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 100
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
//
// x = 22
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
//
// x = 3
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 102
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 103
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 105
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 107
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
// x = 108
// fmt.Printf("%d is at position %d\n", x, jump.Search(arr, x))
//
//}
