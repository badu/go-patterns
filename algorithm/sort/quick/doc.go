package quick

// =====================================================================================================================
// Quick Sort
// =====================================================================================================================
// Like Merge Sort, QuickSort is a Divide and Conquer algorithm.
// It picks an element as pivot and partitions the given array around the picked pivot.
// There are many different versions of quickSort that pick pivot in different ways.
// 1. Always pick first element as pivot.
// 2. Always pick last element as pivot (implemented below)
// 3. Pick a random element as pivot.
// 4. Pick median as pivot.
// The key process in quickSort is partition().
// Target of partitions is, given an array and an element x of array as pivot,
// put x at its correct position in sorted array and put all smaller elements (smaller than x) before x,
// and put all greater elements (greater than x) after x.
// All this should be done in linear time.
//
// Pseudocode for Quick Sort
// ===========================================
// Refer image 2
//
// Partition Algorithm
// ===========================================
// There can be many ways to do partition, following pseudo code adopts the method given in CLRS book.
// The logic is simple, we start from the leftmost element and
// keep track of index of smaller (or equal to) elements as i.
// While traversing, if we find a smaller element, we swap current element with arr[i].
// Otherwise we ignore current element.
//
// Pseudocode for Partition
// ===========================================
// Refer image 3
//
// To see the working of the algorithm refer image 4
//
// Time Complexity Analysis of QuickSort
// ===========================================
// Time taken by QuickSort in general can be written as following.
// T(n) = T(k) + T(n-k-1) + theta(n)
// The first two terms are for two recursive calls, the last term is for the partition process.
// k is the number of elements which are smaller than pivot.
//
// The time taken by QuickSort depends upon the input array and partition strategy.
// Following are three cases.
//
// Time Complexity
// ===========================================
// Worst Case:
// ===========================================
// The worst case occurs when the partition process always picks greatest or smallest element as pivot.
// If we consider above partition strategy where last element is always picked as pivot,
// the worst case would occur when the array is already sorted in increasing or decreasing order.
// This condition leads to the case in which the pivot element lies in an extreme end of the sorted array.
// One sub-array is always empty and another sub-array contains n - 1 elements.
// Thus, quicksort is called only on this sub-array.
// The worst case would be theta(n^2).
//
// Best Case
// ===========================================
// The best case occurs when the partition process always picks the middle element as pivot.
// The best case would be theta(n log n)
//
// Average Case
// ===========================================
// To do average case analysis, we need to consider all possible permutation of array and
// calculate time taken by every permutation which doesn’t look easy.
// Average Case Complexity : O(n*log n)
// Space Complexity: O(n*log n)
//
// When to use
// ===========================================
// Although the worst case time complexity of QuickSort is O(n^2) which is more than many other sorting algorithms
// like Merge Sort and Heap Sort, QuickSort is faster in practice,
// because its inner loop can be efficiently implemented on most architectures, and in most real-world data.
// QuickSort can be implemented in different ways by changing the choice of pivot,
// so that the worst case rarely occurs for a given type of data.
// However, merge sort is generally considered better when data is huge and stored in external storage.
//
// Is QuickSort stable?
// ===========================================
// The default implementation is not stable.
// However any sorting algorithm can be made stable by considering indexes as comparison parameter.
//
// Is QuickSort In-place?
// ===========================================
// As per the broad definition of in-place algorithm it qualifies as an in-place sorting algorithm as it
// uses extra space only for storing recursive function calls but not for manipulating the input.
//
// What is 3-Way QuickSort?
// ===========================================
// In simple QuickSort algorithm, we select an element as pivot,
// partition the array around pivot and recur for subarrays on left and right of pivot.
// Consider an array which has many redundant elements.
// For example, {1, 4, 2, 4, 2, 4, 1, 2, 4, 1, 2, 2, 2, 2, 4, 1, 4, 4, 4}.
// If 4 is picked as pivot in Simple QuickSort, we fix only one 4 and recursively process remaining occurrences.
// In 3 Way QuickSort, an array arr[l..r] is divided in 3 parts:
// a) arr[l..i] elements less than pivot.
// b) arr[i+1..j-1] elements equal to pivot.
// c) arr[j..r] elements greater than pivot.
//
// Why Quick Sort is preferred over MergeSort for sorting Arrays
// ===========================================
// Quick Sort in its general form is an in-place sort (i.e. it doesn’t require any extra storage) whereas
// merge sort requires O(N) extra storage, N denoting the array size which may be quite expensive.
// Allocating and de-allocating the extra space used for merge sort increases the running time of the algorithm.
// Comparing average complexity we find that both type of sorts have O(N log N) average complexity but
// the constants differ. For arrays, merge sort loses due to the use of extra O(N) storage space.
//
// Most practical implementations of Quick Sort use randomized version.
// The randomized version has expected time complexity of O(n Log n).
// The worst case is possible in randomized version also,
// but worst case doesn’t occur for a particular pattern (like sorted array) and
// randomized Quick Sort works well in practice.
//
// Quick Sort is also a cache friendly sorting algorithm as it has good locality of reference when used for arrays.
//
// Quick Sort is also tail recursive, therefore tail call optimizations is done.
//
// Why MergeSort is preferred over QuickSort for Linked Lists?
// ===========================================
// In case of linked lists the case is different mainly due to
// difference in memory allocation of arrays and linked lists.
// Unlike arrays, linked list nodes may not be adjacent in memory.
// Unlike array, in linked list, we can insert items in the middle in O(1) extra space and O(1) time.
// Therefore merge operation of merge sort can be implemented without extra space for linked lists.
//
// In arrays, we can do random access as elements are continuous in memory.
// Let us say we have an integer (4-byte) array A and let the address of A[0] be x then to access A[i],
// we can directly access the memory at (x + i*4).
// Unlike arrays, we can not do random access in linked list.
// Quick Sort requires a lot of this kind of access.
// In linked list to access i’th index,
// we have to travel each and every node from the head to i’th node as we don’t have continuous block of memory.
// Therefore, the overhead increases for quick sort.
// Merge sort accesses data sequentially and the need of random access is low.
//
// How to optimize QuickSort so that it takes O(Log n) extra space in worst case?
// ===========================================
// In QuickSort, partition function is in-place, but we need extra space for recursive function calls.
// A simple implementation of QuickSort makes two calls to itself and in worst case
// requires O(n) space on function call stack.
// The worst case happens when the selected pivot always divides the array such that one part has 0 elements and
// other part has n-1 elements.
// For example, in below code, if we choose last element as pivot,
// we get worst case for sorted arrays (See this for visualization)
//
// Can we reduce the auxiliary space for function call stack?
// ===========================================
// We can limit the auxiliary space to O(Log n).
// The idea is based on tail call elimination.
// As seen in the previous post, we can convert the code so that it makes one recursive call.
// For example, in the below code,
// we have converted the above code to use a while loop and have reduced the number of recursive calls.
// /* QuickSort after tail call elimination using while loop */
//void quickSort(int arr[], int low, int high)
//{
//    while (low < high)
//    {
//        /* pi is partitioning index, arr[p] is now
//           at right place */
//        int pi = partition(arr, low, high);
//
//        // Separately sort elements before
//        // partition and after partition
//        quickSort(arr, low, pi - 1);
//
//        low = pi+1;
//    }
//}
//
// Although we have reduced number of recursive calls, the above code can still use O(n) auxiliary space in worst case.
// In worst case, it is possible that array is divided in a way that the first part always has n-1 elements.
// For example, this may happen when last element is chosen as pivot and array is sorted in decreasing order.
//
// We can optimize the above code to make a recursive call only for the smaller part after partition.
// Below is implementation of this idea.
// /* This QuickSort requires O(Log n) auxiliary space in
//   worst case. */
//void quickSort(int arr[], int low, int high)
//{
//    while (low < high)
//    {
//        /* pi is partitioning index, arr[p] is now
//           at right place */
//        int pi = partition(arr, low, high);
//
//        // If left part is smaller, then recur for left
//        // part and handle right part iteratively
//        if (pi - low < high - pi)
//        {
//            quickSort(arr, low, pi - 1);
//            low = pi + 1;
//        }
//
//        // Else recur for right part
//        else
//        {
//            quickSort(arr, pi + 1, high);
//            high = pi - 1;
//        }
//    }
//}
//
// In the above code, if left part becomes smaller, then we make recursive call for left part.
// Else for the right part.
// In worst case (for space), when both parts are of equal sizes in all recursive calls, we use O(Log n) extra space.
//
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// arr := []int{10, 80, 30, 90, 40, 50, 70}
// quick.Sort(arr, 0, 6)
//
//	fmt.Println("Result: ", arr)
//
//}
//
// Output
// ===========================================
// Result: [10 30 40 50 70 80 90]
