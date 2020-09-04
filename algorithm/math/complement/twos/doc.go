package twos

// =====================================================================================================================
// 2's Complement
// =====================================================================================================================
// 2’s complement of a binary number is 1 added to the 1’s complement of the binary number.
// For one’s complement, we simply need to flip all bits.
// For 2’s complement, we first find one’s complement.
// We traverse the one’s complement starting from LSB (least significant bit),
// and look for 0. We flip all 1’s (change to 0) until we find a 0.
// Finally, we flip the found 0.
// For example, 2’s complement of “01000” is “11000” (Note that we first find one’s complement of 01000 as 10111).
// If there are all 1’s (in one’s complement), we add an extra 1 in the string.
// For example, 2’s complement of “000” is “1000” (1’s complement of “000” is “111”).
//
// Examples
// ===========================================
// 2's complement of "0111" is  "1001"
// 2's complement of "1100" is  "0100"
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// n := "1100"
// fmt.Printf("Complement of %s is %s\n", n, twos.Complement(n))
//
//}
// Output
// ===========================================
// Complement of 1100 is 0100
