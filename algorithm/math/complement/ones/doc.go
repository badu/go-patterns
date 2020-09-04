package ones

// =====================================================================================================================
// 1's Complement
// =====================================================================================================================
// 1’s complement of a binary number is another binary number obtained by toggling all bits
// in it, i.e., transforming the 0 bit to 1 and the 1 bit to 0.
// Golang doesn’t have any specified unary Bitwise NOT(~) or you can say Bitwise Complement operator like
// other programming languages(C/C++, Java, Python, etc).
// Here, you have to use Bitwise XOR(^) operator as Bitwise NOT operator
// Let’s understand how Bitwise XOR takes in any two equal length bit patterns and
// performs Exclusive OR operation on each pair of corresponding bits.
// 1 XOR 1 = 0
// 1 XOR 0 = 1
// 0 XOR 0 = 0
// 0 XOR 1 = 1
// Here, you can see the result of XOR(M, N) = 1 only if M != N else it will be 0.
// So here, we will use the XOR operator as a unary operator to implement the one’s complement to a number.
// Example:
// Input: 11111111 XOR 00001111
// Output: 11110000
//
// Examples
// ===========================================
// 1's complement of "0111" is "1000"
// 1's complement of "1100" is  "0011"
//
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// ones.Complement(15)
//}
// Output
// ===========================================
// Input : 00001111
// Output: 11110000
