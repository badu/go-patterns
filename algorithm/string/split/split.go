package split

// Split splits the given string by the given separator
// Do not use the in built function
// Outputs the same result as strings.Split()
// sLen is the length of the processed string
// tmp is the temporary separated value
// tmp will be reset after a separator is found
// b[i] == sep[0]: Is current value of 'b' == the first character of the sep string
// Do b[i] == sep[0] check only if unprocessed string len is > len(sep)
// s[i:i+lSep]: returns the substring of processed string to len(sep)
// i = i + (lSep - 1): Forwards the processed string by len(sep)
// done when a separator is found
// "-1" since the for loop has i++
// isSep || l == sLen: append if,
// 1. Separator is found
// 2. s is fully processed
func Split(s, sep string) []string {
	b := []byte(s)
	l, lSep := len(s), len(sep)
	r := make([]string, 0, 1)

	var sLen int
	var tmp string
	for i := 0; i < len(b); i++ {
		if (l-sLen >= len(sep)) && (b[i] == sep[0]) {
			isSep := compareSep(s[i:i+lSep], sep)
			if isSep || l == sLen {
				r = append(r, tmp)
				i = i + (lSep - 1)
				sLen++
				tmp = ""
				continue
			}
		}
		sLen++
		tmp = tmp + string(b[i])
	}
	//append the remaining string
	r = append(r, tmp)

	return r
}

//compareSep compares the separator string with the substring passed
func compareSep(sub, sep string) bool {
	if sub == sep {
		return true
	}
	return false
}
