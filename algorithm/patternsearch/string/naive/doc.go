package naive

// =====================================================================================================================
// Naive Pattern Searching
// =====================================================================================================================
// Naive pattern searching is the simplest method among other pattern searching algorithms.
// It checks for all character of the main string to the pattern.
// This algorithm is helpful for smaller texts.
// It does not need any pre-processing phases.
// We can find substring by checking once for the string.
// It also does not occupy extra space to perform the operation.
// Slide the pattern over text one by one and check for a match
// If a match is found, then slide by 1 again to check for subsequent matches
// The Naive pattern searching algorithm doesn’t work well in cases where we see many matching
// characters followed by a mismatching character
//
// Time Complexity
// ===========================================
// The time(worst time) complexity of naive algorithm is O(m(n-m+1)) or ~O(m*n) {for very large inputs)
// The m is the size of pattern and n is the size of the main string.
// Preprocessing: None
//
// When to use
// ===========================================
// 1. Although inefficient, it may be beneficial to use it in cases where the speed advantage of
//    another algorithm is negligible or does not outhweigh the additional setup needed
//    (for example if your source and query pattern are really short)
//
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// text := "AABAACAADAABAAABAA"
// pattern := "AABA"
// fmt.Printf("Pattern found at index %d \n", naive.Search(text, pattern))
//
//}
