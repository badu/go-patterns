package permutation

func Permute(str string, l, r int) []string {
	res := make([]string, 0, 1)

	// Check for str length equal to 1
	// or l has reached the last character of str
	if l == r {
		return []string{str}
	}

	buf := []byte(str)
	for i := l; i <= r; i++ {
		buf[l], buf[i] = buf[i], buf[l]

		// l+1 to fix the lth character of str
		res = append(res, Permute(string(buf), l+1, r)...)

		// Backtrack
		// Restore to the previous state of str after Permute
		buf[l], buf[i] = buf[i], buf[l]
	}
	return res
}
