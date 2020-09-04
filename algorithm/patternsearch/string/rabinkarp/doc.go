package rabinkarp

// =====================================================================================================================
// Rabin Karp Pattern Searching
// =====================================================================================================================
// The Naive String Matching algorithm slides the pattern one by one.
// After each slide, it one by one checks characters at the current shift and
// if all characters match then prints the match.
// Like the Naive Algorithm, Rabin-Karp algorithm also slides the pattern one by one.
// But unlike the Naive algorithm,
// Rabin Karp algorithm matches the hash value of the pattern with the hash value of current substring of text,
// and if the hash values match then only it starts matching individual characters.
// So Rabin Karp algorithm needs to calculate hash values for following strings.
// 1. Pattern itself.
// 2. All the substrings of text of length m.
// Since we need to efficiently calculate hash values for all the substrings of size m of text,
// we must have a hash function which has following property:
// Hash at the next shift must be efficiently computable from the current hash value and
// next character in text or we can say hash(txt[s+1 .. s+m]) must be efficiently computable
// from hash(txt[s .. s+m-1]) and
// txt[s+m] i.e., hash(txt[s+1 .. s+m])= rehash(txt[s+m], hash(txt[s .. s+m-1])) and rehash must be O(1) operation.
// The hash function suggested by Rabin and Karp calculates an integer value.
// The integer value for a string is numeric value of a string.
// For example, if all possible characters are from 1 to 10, the numeric value of “122” will be 122.
// The number of possible characters is higher than 10 (256 in general) and pattern length can be large.
// So the numeric values cannot be practically stored as an integer.
// Therefore, the numeric value is calculated using modular arithmetic to make sure that
// the hash values can be stored in an integer variable (can fit in memory words).
// To do rehashing,
// we need to take off the most significant digit and add the new least significant digit for in hash value.
// Rehashing is done using the following formula:
// hash( txt[s+1 .. s+m] ) = ( d ( hash( txt[s .. s+m-1]) - txt[s]*h ) + txt[s + m] ) mod q
// hash( txt[s .. s+m-1] ) : Hash value at shift s.
// hash( txt[s+1 .. s+m] ) : Hash value at next shift (or shift s+1)
// d: Number of characters in the alphabet
// q: A prime number
// h: d^(m-1)
//
// How does above expression work?
// ===========================================
// This is simple mathematics, we compute decimal value of current window from previous window.
// For example pattern length is 3 and string is “23456”
// You compute the value of first window (which is “234”) as 234.
// How how will you compute value of next window “345”? You will do (234-2*100)*10 + 5 and get 345.
//
// Time Complexity
// ===========================================
// The time(worst time) complexity of Rabin Karp algorithm is O(mn). This happens when,
// every pattern hash matches with text hash but substrings dont match
// The running time of RABIN-KARP-MATCHER in the worst case scenario O ((n-m+1) m )
// but it has a good average case running time.
// If the expected number of strong shifts is small O (1) and prime q is chosen to be quite large, then the Rabin-Karp algorithm can be expected to run in time O (n+m) plus the time to require to process spurious hits.
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
