package lps

// =====================================================================================================================
// Longest Palindromic Subsequence
// =====================================================================================================================
// https://youtu.be/TLaGwTnd3HY
// Given a sequence, find the length of the longest palindromic subsequence in it.
// Refer image1
//
// As another example, if the given sequence is “BBABCBCAB”,
// then the output should be 7 as “BABCBAB” is the longest palindromic subsequence in it.
// “BBBBB” and “BBCBB” are also palindromic subsequences of the given sequence, but not the longest ones.
//
// The naive solution for this problem is to generate all subsequences of the given sequence and
// find the longest palindromic subsequence.
// This solution is exponential in term of time complexity.
// Let us see how this problem possesses both important properties of a Dynamic Programming (DP) Problem and
// can efficiently be solved using Dynamic Programming.
//
// 1) Optimal Substructure:
// Let X[0..n-1] be the input sequence of length n and L(0, n-1) be the length of the
// longest palindromic subsequence of X[0..n-1].
//
// If last and first characters of X are same, then L(0, n-1) = L(1, n-2) + 2.
// Else L(0, n-1) = MAX (L(1, n-1), L(0, n-2)).
//
// 2) Overlapping Subproblems
//
// Time Complexity
// ===========================================
// Time Complexity of the above implementation is O(n^2) which is
// much better than the worst-case time complexity of Naive Recursive implementation which is O(n^3).
//
// Refer lps.go for code