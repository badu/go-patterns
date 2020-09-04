package multiple

// =====================================================================================================================
// Check if a Number is Multiple of 3
// =====================================================================================================================
// The very first solution that comes to our mind is the one that we learned in school.
// If sum of digits in a number is multiple of 3 then number is multiple of 3 e.g.,
// for 612 sum of digits is 9 so it’s a multiple of 3. But this solution is not efficient.
// You have to get all decimal digits one by one, add them and then check if sum is multiple of 3.
//
// There is a pattern in binary representation of the number that can be used to find if number is a multiple of 3.
// If difference between count of odd set bits (Bits set at odd positions) and
// even set bits is multiple of 3, so is the number.
//
// Problem
// ===========================================
// Algorithm: Multiple3(n)
// 1. Make n positive if n is negative.
// 2. If number is 0 then return 1
// 3. If number is 1 then return 0
// 4. Initialize: odd_count = 0, even_count = 0
// 5. Loop for n != 0
//    a) If rightmost bit is set then increment odd count.
//    b) Right-shift n by 1 bit
//    c) If rightmost bit is set then increment even count.
//    d) Right-shift n by 1 bit
// 6. return Multiple3(odd_count - even_count)
//
// Examples
// ===========================================
// Example : 23 (00..10111)
// 1) Get count of all set bits at odd positions (For 23 it’s 3).
// 2) Get count of all set bits at even positions (For 23 it’s 1).
// 3) If difference of above two counts is a multiple of 3 then number is also a multiple of 3.
//
// For 23 it’s 2 so 23 is not a multiple of 3
//
// Time Complexity Example 1
// ===========================================
// Time Complexity: O(log n)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// num := 24
//
// if multiple.Multiple3(num) != 0 {
// 	fmt.Println(num, " is multiple of 3")
// } else {
// 	fmt.Println(num, "is not a multiple of 3")
// }
//
//}
// Output
// ===========================================
//
