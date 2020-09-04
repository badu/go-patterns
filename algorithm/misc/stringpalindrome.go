package misc

import (
	"strings"
)

// IsPalindrome returns true if s is a palindrome otherwise returns false
func IsPalindrome(s string) bool {
	// Make case insensitive by converting "s" into lower case
	s = strings.ToLower(s)

	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		if s[i] != s[sLen-1-i] {
			return false
		}
	}
	return true
}

// RUN INSTRUCTIONS (IsPalindrome):
// ===========================================
// s := "refer"
// fmt.Printf("%s is a palindrome? Result: %t\n", s, misc.IsPalindrome(s))
//
// s = "test"
// fmt.Printf("%s is a palindrome? Result: %t\n", s, misc.IsPalindrome(s))
//
// s = "Refer"
// fmt.Printf("%s is a palindrome? Result: %t\n", s, misc.IsPalindrome(s))
//
// Output
// ===========================================
// refer is a palindrome? Result: true
// test is a palindrome? Result: false
// Refer is a palindrome? Result: true
