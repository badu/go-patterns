package cases

// ToLower converts the given leters in a string from upper case to lower case
func ToLower(s string) string {

	buf := make([]byte, 0, len(s))

	// convert string s into a byte stream
	bs := []byte(s)

	for _, val := range bs {
		if val > 64 && val < 91 {
			buf = append(buf, val+32)
			continue
		}
		buf = append(buf, val)
	}
	return string(buf)
}
