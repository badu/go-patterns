package subsetsum

// =====================================================================================================================
// Subset Sum Problem
// =====================================================================================================================
// Given a set of non-negative integers, and a value sum,
// determine if there is a subset of the given set with sum equal to given sum.
//
// https://youtu.be/s6FhG--P7z0
//
// Example
// ===========================================
// Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 9
// Output: True
// There is a subset (4, 5) with sum 9.
//
// Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 30
// Output: False
// There is no subset that add up to 30.
//
//
// Time Complexity
// ===========================================
// Time Complexity is  O(m * n)
// Auxiliary Space: O(m * n)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{2, 3, 7, 8, 10}
// fmt.Println(subsetsum.SubsetSum(arr, 11))
//
// arr = []int{3, 34, 4, 12, 5, 2}
// fmt.Println(subsetsum.SubsetSum(arr, 9))
//
//}
// Output
// ===========================================
// Subset:  [8 3]
// true
// Subset:  [5 4]
// true
