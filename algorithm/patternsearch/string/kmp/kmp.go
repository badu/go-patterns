package kmp

func Search(txt, pat string) []int {

	res := make([]int, 0, 1)

	n := len(txt)
	m := len(pat)

	lps := make([]int, m, m)
	computeLPS(pat, m, lps)

	// Pointer for txt
	i := 0
	// Pointer for pat
	j := 0
	for i < n {
		if txt[i:i+1] == pat[j:j+1] {
			i += 1
			j += 1
		}
		if j == m {
			res = append(res, i-j)
			j = lps[j-1]
		} else if i < n && pat[j:j+1] != txt[i:i+1] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i += 1
			}
		}
	}
	return res
}
func computeLPS(pattern string, m int, lps []int) {
	lpsLen := 0
	i := 1
	lps[0] = 0

	for i < m {
		if pattern[i:i+1] == pattern[lpsLen:lpsLen+1] {
			lps[i] = lpsLen + 1
			lpsLen += 1
			i += 1
		} else {
			if lpsLen != 0 {
				lpsLen = lps[lpsLen-1]
			} else {
				lps[i] = 0
				i += 1
			}
		}
	}
}
