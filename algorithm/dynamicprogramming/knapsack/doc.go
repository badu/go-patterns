package knapsack

// =====================================================================================================================
// 0-1 Knapsack Problem
// =====================================================================================================================
// https://youtu.be/8LusJS5-AGo
// https://youtu.be/nLmhmB6NzcM
//
// Example
// ===========================================
// Let weight elements = {1, 2, 3}
// Let weight values = {10, 15, 40}
// Capacity=6
//
//   0   1   2   3   4   5   6
//
// 0  0   0   0   0   0   0   0
//
// 1  0  10  10  10  10  10  10
//
// 2  0  10  15  25  25  25  25
//
// Explanation:
// For filling 'weight = 2' we come
// across 'j = 3' in which
// we take maximum of
// (10, 15 + DP[1][3-2]) = 25
//   |        |
// '2'       '2 filled'
// not filled
//
// Time Complexity
// ===========================================
// Time Complexity is O(W*CAP)
// Auxiliary Space: O(W*CAP)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// w := []int{2, 3, 4, 5}
// p := []int{1, 2, 5, 6}
// fmt.Println("KnapSack: ", knapsack.Knapsack(w, p, 8))
//
// w = []int{10, 20, 30}
// p = []int{60, 100, 120}
// fmt.Println("KnapSack: ", knapsack.Knapsack(w, p, 50))
//
//}
// Output
// ===========================================
// KnapSack:  8
// KnapSack:  220
