package reverse

//Reverses the given string without using inbuilt functions
func ReverseStringNoInbuiltFunc(str string) string {
	resStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		resStr = resStr + string(str[i])
	}
	return resStr
}
