package kmp

// =====================================================================================================================
// Knuth Morris Pratt(KMP) Pattern Searching
// =====================================================================================================================
// The KMP algorithm makes use of previous match information to determine an amount of skips that
// can be made until the next position in the source gets examined as a possible match.
// To achieve that, a prefix table (or failure function) of the pattern needs to be computed,
// which determines the amount of skippable elements depending on the previous (partial) match.
// The KMP matching algorithm uses degenerating property
// (pattern having same sub-patterns appearing more than once in the pattern)
// of the pattern and improves the worst case complexity to O(n).
// The basic idea behind KMP’s algorithm is:
// whenever we detect a mismatch (after some matches),
// we already know some of the characters in the text of the next window.
// We take advantage of this information to avoid matching the characters that we know will anyway match.
// For explanation refer:
// https://www.youtube.com/watch?v=4jY57Ehc14Y&list=PLZr8BEt9JH1TFPbEq2r1NPM21jyEp3Tmp
//
// Need of Preprocessing?
// ===========================================
// An important question arises from the above explanation,
// how to know how many characters to be skipped. To know this,
// we pre-process pattern and prepare an integer array
// lps[] that tells us the count of characters to be skipped.
//
// Preprocessing Overview:
// ===========================================
// KMP algorithm preprocesses pat[] and constructs an auxiliary lps[] of size m
// (same as size of pattern) which is used to skip characters while matching.
//
// Time Complexity
// ===========================================
// The time(worst time) complexity of KMP algorithm is O(n)
// Preprocessing: Θ(m)
//
// When to use
// ===========================================
// 1.
//
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// txt := "onionionspl"
// pat := "onions"
// fmt.Printf("Pattern found at index %d \n", kmp.Search(txt, pat))
// txt = "ABABDABACDABABCABAB"
// pat = "ABABCABAB"
// fmt.Printf("Pattern found at index %d \n", kmp.Search(txt, pat))
// txt = "AABAACAADAABAAABAA"
// pat = "AABA"
// fmt.Printf("Pattern found at index %d \n", kmp.Search(txt, pat))
//
//}
// Output
// ===========================================
// Pattern found at index [3]
// Pattern found at index [10]
// Pattern found at index [0 9 13]
