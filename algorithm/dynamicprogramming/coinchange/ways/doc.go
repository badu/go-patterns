package ways

// =====================================================================================================================
// Coin Change(No. of ways)
// =====================================================================================================================
// https://youtu.be/L27_JpN6Z1Q (Taken an addition 0 row in addition to the explanation)
//
//
// Example
// ===========================================
// for N = 4 and S = {1,2,3}
// There are four solutions: {1,1,1,1},{1,1,2},{2,2},{1,3}.
// So output should be 4
//
// Time Complexity
// ===========================================
// Time Complexity is O(m*n)
// Auxiliary Space: O(m*n)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// coins := []int{1, 2, 3}
// fmt.Println("Minimum coins: ", ways.CoinChange(coins, 4))
// coins = []int{1, 2, 3}
// fmt.Println("Minimum coins: ", ways.CoinChange(coins, 5))
//
//}
// Output
// ===========================================
// Minimum coins:  4
// Minimum coins:  5
