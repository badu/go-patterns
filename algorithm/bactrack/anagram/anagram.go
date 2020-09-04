package anagram

import (
	"fmt"
)

// This function returns true if contents
// of pMap and tMap are same, otherwise
// false.
func compare(pMap, tMap map[uint8]int) bool {
	for i, pat := range pMap {
		if pat != tMap[i] {
			return false
		}
	}
	return true
}

// Search searches for all permutations of pat in txt
func Search(txt, pat string) {
	m := len(pat)
	n := len(txt)

	// pMap[] stores count of all characters of pattern
	// tMap[] stores count of current window of text
	pMap := make(map[uint8]int)
	tMap := make(map[uint8]int)

	// Calculate pMap and tMap for the first window
	for i := 0; i < m; i++ {
		pMap[pat[i]]++
		tMap[txt[i]]++
	}

	// Traverse through remaining characters of pattern
	for i := m; i < n; i++ {
		// Compare counts of current window
		// of text with counts of pattern[]
		if compare(pMap, tMap) {
			fmt.Println("Found at Index ", i-m)
		}

		// Add current character to current window
		tMap[txt[i]]++

		// Remove the first character of previous window
		tMap[txt[i-m]]--
	}

	// Check for the last window in text
	if compare(pMap, tMap) {
		fmt.Println("Found at Index ", n-m)
	}

}
