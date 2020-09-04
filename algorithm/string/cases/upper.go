package cases

// ToUpper converts the given lower case letters in a string to upper case
func ToUpper(s string) string {

	buf := make([]byte, 0, len(s))

	// convert string s into a byte stream
	bs := []byte(s)

	for _, val := range bs {
		if val > 96 && val < 124 {
			buf = append(buf, val-32)
			continue
		}
		buf = append(buf, val)
	}
	return string(buf)
}
