package ternary

// =====================================================================================================================
// Ternary Search
// =====================================================================================================================
//
// Binary vs Ternary Search
// ===========================================
// Which of the above two does less comparisons in worst case?
// From the first look, it seems the ternary search does less number of comparisons as it makes Log3n recursive calls,
// but binary search makes Log2n recursive calls. Let us take a closer look.
// The following is recursive formula for counting comparisons in worst case of Binary Search.
// T(n) = T(n/2) + 2,  T(1) = 1
//
// The following is recursive formula for counting comparisons in worst case of Ternary Search.
// T(n) = T(n/3) + 4, T(1) = 1
// In binary search, there are 2Log2n + 1 comparisons in worst case.
// In ternary search, there are 4Log3n + 1 comparisons in worst case.
//
// Time Complexity for Binary search = 2clog2n + O(1)
// Time Complexity for Ternary search = 4clog3n + O(1)
// Therefore, the comparison of Ternary and Binary Searches boils down the comparison of expressions 2Log3n and Log2n
// The value of 2Log3n can be written as (2 / Log2 3) * Log2 n .
// Since the value of (2 / Log2 3) is more than one,
// Ternary Search does more comparisons than Binary Search in worst case.
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91, 95, 96, 97, 98, 99, 100, 103, 105, 107}
// x = 2
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 5
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 8
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 12
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 16
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 23
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 38
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 56
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 72
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 91
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 95
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 96
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 97
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 98
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 99
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 100
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
//
// x = 22
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
//
// x = 3
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 102
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 103
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 105
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 107
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
// x = 108
// fmt.Printf("%d is at position %d\n", x,  ternary.Search(arr, x))
//
//}
//
// Output
// ===========================================
// 2 is at position 0
// 5 is at position 1
// 8 is at position 2
// 12 is at position 3
// 16 is at position 4
// 23 is at position 5
// 38 is at position 6
// 56 is at position 7
// 72 is at position 8
// 91 is at position 9
// 95 is at position 10
// 96 is at position 11
// 97 is at position 12
// 98 is at position 13
// 99 is at position 14
// 100 is at position 15
// 22 is at position -1
// 3 is at position -1
// 101 is at position -1
// 101 is at position -1
// 102 is at position -1
// 103 is at position 16
// 105 is at position 17
// 107 is at position 18
// 108 is at position -1
