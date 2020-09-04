package firstunique

func FirstUnique(str string) string {
	sLen := len(str)
	charCount := make([]int, 256, 256)

	// Take character ascii value as index to charCount
	// Increment count when a char is found
	for i := 0; i < sLen; i++ {
		charCount[str[i]]++
	}

	// The first occurrence of charCount with value == 1 is
	// the first unique character in the string str
	var res int
	for i := 0; i < len(charCount); i++ {
		if charCount[i] == 1 {
			res = i
			break
		}
	}
	return string(res)
}
