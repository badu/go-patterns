package lcs

// =====================================================================================================================
// Longest Common Subsequence
// =====================================================================================================================
// https://youtu.be/sSno9rV8Rhg
// https://youtu.be/NnD96abizww
// A subsequence is a sequence that appears in the same relative order, but not necessarily contiguous
// For example, “abc”, “abg”, “bdf”, “aeg”, ‘”acefg”, .. etc are subsequences of “abcdefg”.
//
// Example
// ===========================================
// LCS for input Sequences “ABCDGH” and “AEDFHR” is “ADH” of length 3.
// LCS for input Sequences “AGGTAB” and “GXTXAYB” is “GTAB” of length 4.
//
// Time Complexity
// ===========================================
// Time complexity is O(m*n)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// s1 := "bd"
// s2:= "abcd"
// fmt.Println("LCS: ", lcs.LCS(s1, s2))
//
// s1 = "AGGTAB"
// s2 = "GXTXAYB"
// fmt.Println("LCS: ", lcs.LCS(s1, s2))
//
//}
// Output
// ===========================================
// LCS:  2
// LCS:  4