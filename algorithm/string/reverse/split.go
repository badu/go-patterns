package reverse

import (
	"strings"
)

//Solve using string.Split(str, sep) function
//The range for loop syntax can also be used to iterate over a string
//When used it returns the ascii values of each letter
func ReverseStringUseSplit(str string) string {
	resStr := ""
	tmpStrArr := strings.Split(str, "")
	for i := len(tmpStrArr) - 1; i >= 0; i-- {
		resStr = resStr + tmpStrArr[i]
	}
	return resStr
}
