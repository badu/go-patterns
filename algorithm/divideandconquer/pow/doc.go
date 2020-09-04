package pow

// =====================================================================================================================
// Write a program to calculate pow(x,n)
// =====================================================================================================================
//
// Problem
// ===========================================
// Given two integers x and n, write a function to compute xn.
// We may assume that x and n are small and overflow doesn’t happen.
//
// Solution in example 1 can be optimized to O(log n) by calculating power(x, y/2) only once and storing it.
//
// Examples
// ===========================================
// Input : x = 2, n = 3
// Output : 8
//
// Input : x = 7, n = 2
// Output : 49
//
// Time Complexity Example 1
// ===========================================
// Time Complexity: O(n)
// Space Complexity: O(1)
// Algorithmic Paradigm: Divide and conquer.
//
// Time Complexity Example 2
// ===========================================
// Time Complexity: O(log n)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// fmt.Println(e1.Pow(2, 5))
// fmt.Println(e1.PowRecur(2, 5))
// fmt.Println(e2.Pow(2, 5))
//
//}
// Output
// ===========================================
// 32
// 32
// 32
