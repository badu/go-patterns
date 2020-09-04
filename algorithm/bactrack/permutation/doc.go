package permutation

// =====================================================================================================================
// Permutation of a String
// =====================================================================================================================
// A permutation, also called an “arrangement number” or “order,”
// is a rearrangement of the elements of an ordered list S into a one-to-one correspondence with S itself.
// A string of length n has n! permutation.
// Refer:
// https://www.youtube.com/watch?v=GuTPwotSdYw&list=PLZr8BEt9JH1TFPbEq2r1NPM21jyEp3Tmp&index=4
//
//
// Examples
// ===========================================
// Below are the permutations of string ABC.
// ABC ACB BAC BCA CBA CAB
//
//
//
// Time Complexity
// ===========================================
// Time complexity of this algorithm is O(n!)
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// str := "ABC"
// p := permutation.Permute(str, 0, len(str)-1)
//
// fmt.Println("Permutations: ", p)
//}
// Output
// ===========================================
// Permutations:  [ABC ACB BAC BCA CBA CAB]
