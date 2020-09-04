package heap

// =====================================================================================================================
// Heap Sort
// =====================================================================================================================
// https://youtu.be/HqPJF2L5h9U
//
// What is Heap Data Structure?
// ===========================================
// Heap is a special tree-based data structure.
// A binary tree is said to follow a heap data structure if
// 1. it is a complete binary tree
// 2. All nodes in the tree follow the property that they are greater than their children
//    i.e. the largest element is at the root and both its children and smaller than the root and so on.
//    Such a heap is called a max-heap. If instead, all nodes are smaller than their children, it is called a min-heap
//
// How to "heapify" a tree
// ===========================================
// Starting from a complete binary tree,
// we can modify it to become a Max-Heap by running a function called heapify on all the non-leaf elements of the heap.
// To build a max-heap from any tree, we can thus start heapifying each sub-tree from the bottom up and
// end up with a max-heap after the function is applied to all the elements including the root element.
// In the case of a complete tree, the first index of a non-leaf node is given by n/2 - 1.
// All other nodes after that are leaf-nodes and thus don't need to be heapified.
//
// Heap sort is a comparison based sorting technique based on Binary Heap data structure.
// It is similar to selection sort where we first find the maximum element and place the maximum element at the end.
// We repeat the same process for remaining element.
//
// A Binary Heap is a Complete Binary Tree where items are stored in a special order such that
// value in a parent node is greater(or smaller) than the values in its two children nodes.
// The former is called as max heap and the latter is called min heap.
// The heap can be represented by binary tree or array.
//
// Why array based representation for Binary Heap?
// ===========================================
// Since a Binary Heap is a Complete Binary Tree,
// it can be easily represented as array and array based representation is space efficient.
// If the parent node is stored at index I,
// the left child can be calculated by 2 * I + 1 and right child by 2 * I + 2 (assuming the indexing starts at 0).
//
// Heap Sort Algorithm for sorting in increasing order:
// ===========================================
// 1. Build a max heap from the input data.
// 2. At this point,
//    the largest item is stored at the root of the heap.
//    Replace it with the last item of the heap followed by reducing the size of heap by 1.
//    Finally, heapify the root of tree.
// 3. Repeat above steps while size of heap is greater than 1.
//
// How to build the heap?
// ===========================================
// Heapify procedure can be applied to a node only if its children nodes are heapified.
// So the heapification must be performed in the bottom up order.
//
// Notes:
// ===========================================
// Heap sort is an in-place algorithm.
// Its typical implementation is not stable, but can be made stable
// Heap sort algorithm has limited uses because Quicksort and Mergesort are better in practice.
// Nevertheless, the Heap data structure itself is enormously used.
//
// Example
// ===========================================
// Lets understand with the help of an example:
// Input data: 4, 10, 3, 5, 1
//         4(0)
//        /   \
//     10(1)   3(2)
//    /   \
// 5(3)    1(4)
//
//The numbers in bracket represent the indices in the array
//representation of data.
//
//Applying heapify procedure to index 1:
//         4(0)
//        /   \
//    10(1)    3(2)
//    /   \
//5(3)    1(4)
//
//Applying heapify procedure to index 0:
//        10(0)
//        /  \
//     5(1)  3(2)
//    /   \
// 4(3)    1(4)
//The heapify procedure calls itself recursively to build heap
// in top down manner.
//
// Time Complexity
// ===========================================
// Time complexity of heapify is O(Log n).
// Time complexity of createAndBuildHeap() is O(n) and overall time complexity of Heap Sort is O(n Log n).
// Worst Case Time Complexity: O(n*log n)
// Best Case Time Complexity: O(n*log n)
// Average Time Complexity: O(n*log n)
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// arr := []int{10, 80, 30, 90, 40, 50, 70}
// heap.Sort(arr)
//
// fmt.Println("Result: ", arr)
//
//}
//
// Output
// ===========================================
// Result: [10 30 40 50 70 80 90]
