package misc

import (
	"strconv"
	"strings"
)

//https://practice.geeksforgeeks.org/problems/occurences-of-2-as-a-digit/1
func CountOccurrencesOfTwo(n int) int {
	countTwo := 0

	tmpStr := ""
	for i := 1; i <= n; i++ {
		//strings(i) wont work here
		tmpStr = strconv.Itoa(i)
		strSplit := strings.Split(tmpStr, "")
		for _, val := range strSplit {
			if val[0] == 50 {
				countTwo++
			}
		}
	}
	return countTwo
}
