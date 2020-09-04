package v1

// =====================================================================================================================
// Fibonacci Numbers(DP Space Optimized Method 2)
// =====================================================================================================================
// We can optimize the space used in method 2 by storing the previous two numbers only because that is
// all we need to get the next Fibonacci number in series.
//
//
//
// Time Complexity
// ===========================================
// Time Complexity:O(n)
// Extra Space: O(1)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// fmt.Println(v1.Fibo(5))
//
//}
// Output
// ===========================================
// 5
