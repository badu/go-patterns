package editdistance

// =====================================================================================================================
// Minimum Edit Distance
// =====================================================================================================================
// https://youtu.be/We3YDTzNXEk
// Given two strings str1 and str2 and below operations that can performed on str1. Find minimum number of edits (operations) required to convert ‘str1’ into ‘str2’.
//
// Insert
// Remove
// Replace
// All of the above operations are of equal cost.
//
// Example
// ===========================================
// Input:   str1 = "geek", str2 = "gesek"
// Output:  1
// We can convert str1 into str2 by inserting a 's'.
//
// Input:   str1 = "cat", str2 = "cut"
// Output:  1
// We can convert str1 into str2 by replacing 'a' with 'u'.
//
// Time Complexity
// ===========================================
// Time Complexity is  O(m * n)
// Auxiliary Space: O(m * n)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// s1 := "azced"
// s2 := "abcdef"
// fmt.Println("Minimum Edits: ", editdistance.MinimumEdits(s1, s2))
//
// s1 = "sunday"
// s2 = "saturday"
// fmt.Println("Minimum Edits: ", editdistance.MinimumEdits(s1, s2))
//
//}
// Output
// ===========================================
// Minimum Edits:  3
// Minimum Edits:  3
