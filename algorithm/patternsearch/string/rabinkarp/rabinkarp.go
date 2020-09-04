package rabinkarp

func Search(txt, pat string) []int {
	res := make([]int, 0, 1)

	n := len(txt)
	m := len(pat)

	// Base
	d := 10
	var i, j int
	// Hash value of the text
	var t int
	// Hash value of the pattern
	var p int
	h := 1 // 10^0
	// Value to take mod
	// Select a prime number
	q := 101

	// The value of h would be "pow(d, M-1)%q"
	for i = 0; i < m-1; i++ {
		h = (h * d) % q
	}

	// Calculate the hash value of pattern and first
	// window of text
	for i = 0; i < m; i++ {
		p = (d*p + int(pat[i])) % q
		t = (d*t + int(txt[i])) % q
	}

	for i = 0; i <= n-m; i++ {
		// Check the hash values of current window of text
		// and pattern. If the hash values match then only
		// check for characters on by one
		if p == t {
			for j = 0; j < m; j++ {
				if txt[i+j:i+j+1] != pat[j:j+1] {
					break

				}
			}

			// if p == t and pat[0...M-1] = txt[i, i+1, ...i+M-1]
			if j == m {
				res = append(res, i)
			}
		}

		// Calculate hash value for next window of text: Remove
		// leading digit, add trailing digit
		if i < n-m {
			t = (d*(t-int(txt[i])*h) + int(txt[i+m])) % q

			// We might get negative value of t, converting it
			// to positive
			if t < 0 {
				t = t + q
			}
		}
	}

	return res
}
