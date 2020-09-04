package insertion

// =====================================================================================================================
// Insertion Sort
// =====================================================================================================================
// Insertion sort works similarly as we sort cards in our hand in a card game.
// We assume that the first card is already sorted then, we select an unsorted card.
// If the unsorted card is greater than the card in hand, it is placed on the right otherwise, to the left.
// In the same way, other unsorted cards are taken and put at their right place.
// A similar approach is used by insertion sort.
// Insertion sort is a sorting algorithm that places an unsorted element at its suitable place in each iteration.
//
// What is Binary Insertion Sort?
// ===========================================
// We can use binary search to reduce the number of comparisons in normal insertion sort.
// Binary Insertion Sort uses binary search to find the proper location to insert the selected item at each iteration.
// In normal insertion, sorting takes O(i) (at ith iteration) in worst case.
// We can reduce it to O(log i) by using binary search.
// The algorithm, as a whole,
// still has a running worst case running time of O(n^2) because of the series of swaps required for each insertion.
//
// How to implement Insertion Sort for Linked List?
// ===========================================
// Below is simple insertion sort algorithm for linked list.
//
// 1) Create an empty sorted (or result) list
// 2) Traverse the given list, do following for every node.
// ......a) Insert current node in sorted way in sorted or result list.
// 3) Change head of given linked list to head of sorted (or result) list.
//
// Following are some of the important characteristics of Insertion Sort:
// ===========================================
// 1. It is efficient for smaller data sets, but very inefficient for larger lists.
// 2. Insertion Sort is adaptive, that means it reduces its total number of steps if a partially sorted array is
//    provided as input, making it efficient.
// 3. It is better than Selection Sort and Bubble Sort algorithms.
// 4. Its space complexity is less. Like bubble Sort, insertion sort also requires a single additional memory space.
// 5. It is a stable sorting technique, as it does not change the relative order of elements which are equal.
//
// Algorithm
// ===========================================
// 1. We start by making the second element of the given array, i.e. element at index 1, the key.
// 2. We compare the key element with the element(s) before it, in this case, element at index 0
//    a. If the key element is less than the first element, we insert the key element before the first element.
//    b. If the key element is greater than the first element, then we insert it after the first element.
// 3. Then, we make the third element of the array as key and
//    will compare it with elements to it's left and insert it at the right position.
// 4. And we go on repeating this, until the array is sorted.
//
// Example
// ===========================================
// Input: 12, 11, 13, 5, 6
//
// Let us loop for i = 1 (second element of the array) to 4 (last element of the array)
// i = 1. Since 11 is smaller than 12, move 12 and insert 11 before 12
// 11, 12, 13, 5, 6
//
// i = 2. 13 will remain at its position as all elements in A[0..I-1] are smaller than 13
// 11, 12, 13, 5, 6
//
// i = 3. 5 will move to the beginning and all other elements from 11 to 13 will move one position ahead of their current position.
// 5, 11, 12, 13, 6
//
// i = 4. 6 will move to position after 5, and elements from 11 to 13 will move one position ahead of their current position.
// 5, 6, 11, 12, 13
//
// Time Complexity
// ===========================================
// Worst Case Time Complexity [ Big-O ]: O(n^2)
// Best Case Time Complexity [Big-omega]: O(n)
// Average Time Complexity [Big-theta]: O(n^2)
// Auxiliary Space: O(1)
//
// Boundary Cases: Insertion sort takes maximum time to sort if elements are sorted in reverse order.
// And it takes minimum time (Order of n) when elements are already sorted.
// Algorithmic Paradigm: Incremental Approach
// Sorting In Place: Yes
// Stable: Yes
// Online: Yes
// Uses: Insertion sort is used when number of elements is small.
// It can also be useful when input array is almost sorted, only few elements are misplaced in complete big array.
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// arr := []int{95, 80, 30, 90, 50, 40, 70, 10}
// insertion.Sort(arr)
// fmt.Println("Result: ", arr)
//}
//
// Output
// ===========================================
// Result:  [10 30 40 50 70 80 90 95]
