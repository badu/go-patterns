package docs

// =====================================================================================================================
// Search Algorithms Questions-1
// =====================================================================================================================
//
// What is an algorithm?
// ===========================================
// Informally, an algorithm is any well-defined computational procedure that takes some value, or set of values,
// as input and produces some value, or set of values, as output.
// An algorithm is thus a sequence of computational steps that transform the input into the output.
//
// What is time complexity of Binary Search?
// ===========================================
// Time complexity of binary search is O(Log n)
//
// Can Binary Search be used for linked lists?
// ===========================================
// It's possible, but since random access is not allowed in linked list,
// we cannot reach the middle element in O(1) time.
// Therefore Binary Search for linked lists is inefficient.
//
// Yes, Binary search is possible on the linked list if the list is ordered and you know the count of elements in list.
// But While sorting  the list, you can access a single element at a time through a pointer to that node
// i.e. either a previous node or next node.
// This increases the traversal steps per element in linked list just to find the middle element.
// This makes it slow and inefficient.
//
// Whereas the binary search on an array is fast and efficient because of
// its ability to access any element in the array in constant time.
// You can get to the middle of the array just by saying "array[middle]"!
// Now, can you do the same with a linked list? The answer is No.
// You will have to write your own,
// possibly inefficient algorithm to get the value of the middle node of a linked list.
// In a linked list, you loose the ability to get the value of any node in a constant time.
