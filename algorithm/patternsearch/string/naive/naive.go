package naive

func Search(text, pattern string) []int {
	res := make([]int, 0, 1)

	tLen := len(text)
	pLen := len(pattern)

	for i := 0; i <= tLen-pLen; i++ {
		var j int
		for j = 0; j < pLen; j++ {
			if text[i+j:i+j+1] != pattern[j:j+1] {
				break
			}
		}
		if j == pLen {
			res = append(res, i)
		}
	}
	return res
}
