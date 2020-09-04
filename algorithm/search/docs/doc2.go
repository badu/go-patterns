package docs

// =====================================================================================================================
// Linear Search vs Binary Search
// =====================================================================================================================
// A linear search scans one item at a time, without jumping to any item .
// 1. The worst case complexity is  O(n), sometimes known an O(n) search
// 2. Time taken to search elements keep increasing as the number of elements are increased.
//
// A binary search however, cuts down your search to half as soon as you find middle of a sorted list.
// 1. The middle element is looked to check if it is greater than or less than the value to be searched.
// 2. Accordingly, search is done to either half of the given list
//
// Important Differences
// ===========================================
// Input data needs to be sorted in Binary Search and not in Linear Search
// Linear search does the sequential access whereas Binary search access data randomly.
// Time complexity of linear search is O(n) , Binary search has time complexity O(log n).
// Linear search performs equality comparisons and Binary search performs ordering comparisons
