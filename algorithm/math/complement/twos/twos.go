package twos

func Complement(n string) string {

	buf := []byte(n)

	// Take 1's complement
	for i, bit := range buf {
		if bit == 48 {
			buf[i] = 49
			continue
		}
		buf[i] = 48
	}

	// Replace all LSB 1's with 0
	// Replace the first LSB 0 with 1
	for i := len(buf) - 1; i >= 0; i-- {
		if buf[i] == 49 {
			buf[i] = 48
			continue
		}
		buf[i] = 49
		break
	}
	return string(buf)
}
