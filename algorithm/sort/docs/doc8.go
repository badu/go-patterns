package docs

// =====================================================================================================================
// When to use each Sorting Algorithm
// =====================================================================================================================
// A sorting algorithm is an algorithm that makes arrange in a certain order.
// The fundamental task is to put the items in the desired order so that
// the records are re-arranged for making searching easier.
//
// Below is one by one description for when to use which sorting algorithms for better performance:
// ===========================================
//
// Selection Sort
// ===========================================
// This sorting algorithm sorts an array by repeatedly finding the minimum element
// (considering ascending order) from the unsorted part and putting it at the beginning.
// The algorithm maintains two sub-arrays in a given array,
// the subarray which is already sorted and remaining subarray which is unsorted.
// In every iteration of selection sort, the minimum element (considering ascending order) from
// the unsorted subarray is picked and moved to the sorted subarray.
//
// We can use Selection Sort as per below constraints :
// 1. When the list is small.
//    As the time complexity of selection sort is O(n^2) which makes it inefficient for a large list.
// 2. When memory space is limited because it makes the minimum possible number of swaps during sorting.
//
// Bubble Sort
// ===========================================
// This sorting algorithm is the simplest sorting algorithm that works by
// repeatedly swapping the adjacent elements if they are in the wrong order.
// If we have total N elements, then we need to repeat the above process for N-1 times.
//
// We can use Bubble Sort as per below constraints :
// 1. It works well with large datasets where the items are almost sorted because
//    it takes only one iteration to detect whether the list is sorted or not.
//    But if the list is unsorted to a large extent then this algorithm holds good for small datasets or lists.
//    This algorithm is fastest on an extremely small or nearly sorted set of data.
//
// Insertion Sort
// ===========================================
// This sorting algorithm is a simple sorting algorithm that works the way we sort playing cards in our hands.
// It places an unsorted element at its suitable place in each iteration.
//
// We can use Insertion Sort as per below constraints :
// 1. If the data is nearly sorted or when the list is small as it has a complexity of O(N^2) and
//    if the list is sorted a minimum number of elements will slide over to insert the element at it’s correct location.
// 2. This algorithm is stable and it has fast running case when the list is nearly sorted.
// 3. When usage of memory is a constraint as it has space complexity of O(1).
//
// Merge Sort
// ===========================================
// This sorting algorithm is based on Divide and Conquer algorithm.
// It divides input array into two halves, calls itself for the two halves, and then merges the two sorted halves.
// The merge() function is used for merging two halves.
// The merge(arr, l, m, r) is a key process that assumes that arr[l..m] and arr[m+1..r] are sorted and
// merges the two sorted sub-arrays into one.
//
// We can use Merge Sort as per below constraints :
// 1. Merge sort is used when the data structure doesn’t support random access since
//    it works with pure sequential access that is forward iterators, rather than random access iterators.
// 2. It is widely used for external sorting, where random access can be very,
//    very expensive compared to sequential access.
// 3. It is used where it is known that the data is similar data.
// 4. Merge sort is fast in the case of a linked list.
// 5. It is used in the case of a linked list as in linked list for accessing any data at
//    some index we need to traverse from the head to that index and
//    merge sort accesses data sequentially and the need of random access is low.
// 6. The main advantage of the merge sort is its stability, the elements compared equally retain their original order.
//
// Quick Sort
// ===========================================
// This sorting algorithm is also based on Divide and Conquer algorithm.
// It picks an element as pivot and partitions the given list around the picked pivot.
// After partitioning the list on the basis of the pivot element,
// the Quick is again applied recursively to two sub-lists
// i.e., sublist to the left of the pivot element and sublist to the right of the pivot element.
//
// We can use Quick Sort as per below constraints :
// 1. Quick sort is fastest, but it is not always O(N*log N), as there are worst cases where it becomes O(N^2).
// 2. Quicksort is probably more effective for datasets that fit in memory.
//    For larger data sets it proves to be inefficient so algorithms like merge sort are preferred in that case.
// 3. Quick Sort in is an in-place sort (i.e. it doesn’t require any extra storage) so
//    it is appropriate to use it for arrays.
