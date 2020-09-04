package lis

// =====================================================================================================================
// Longest Increasing Subsequence
// =====================================================================================================================
// https://youtu.be/CE2b_-XfVDk
// The Longest Increasing Subsequence (LIS) problem is to find the length of the
// longest subsequence of a given sequence such that all elements of the subsequence are sorted in increasing order
// For example, the length of LIS for {10, 22, 9, 33, 21, 50, 41, 60, 80} is 6 and LIS is {10, 22, 33, 50, 60, 80}.
//
// Example
// ===========================================
// Input: arr[] = {3, 10, 2, 1, 20}
// Output: Length of LIS = 3
// The longest increasing subsequence is 3, 10, 20
//
// Input: arr[] = {3, 2}
// Output: Length of LIS = 1
// The longest increasing subsequences are {3} and {2}
//
// Input: arr[] = {50, 3, 10, 7, 40, 80}
// Output: Length of LIS = 4
// The longest increasing subsequence is {3, 7, 40, 80}
//
// Time Complexity
// ===========================================
// O(n^2) as nested loop is used
// Auxiliary Space: O(n) as we use an array to store LIS values at each index.
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr:= []int{10, 22, 9, 33, 21, 50, 41, 60}
// fmt.Println(lis.LIS(arr))
//}
// Output
// ===========================================
// 5
//
// RUN INSTRUCTIONS 2:
// ===========================================
// func main() {
// arr:= []int{10, 22, 9, 33, 21, 50, 41, 60}
// fmt.Println(lis.LISSequence(arr))
//}
// Output
// ===========================================
// 5 [10 22 33 50 60]
