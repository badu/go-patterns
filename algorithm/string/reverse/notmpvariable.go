package reverse

func NoTmpVar(str string) string {
	sLen := len(str)
	for i := sLen - 1; i >= 0; i-- {
		str = str + str[i:i+1]
	}
	str = str[len(str)-sLen:]

	return str
}
