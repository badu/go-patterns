package misc

import (
	"api/algorithm/patternsearch/string/rabinkarp"
)

// RemoveChar removes all character/sequence occurrences of "seq" from a given string "str"
func RemoveChar(str, seq string) string {

	var res string
	occArr := rabinkarp.Search(str, seq)
	m := convertArrToMap(occArr)

	seqLen := len(seq)
	for i := 0; i < len(str); {
		if m[i] {
			i = i + seqLen
			continue
		}
		res = res + str[i:i+1]
		i++
	}
	return res
}

func convertArrToMap(arr []int) map[int]bool {
	res := make(map[int]bool)
	for _, val := range arr {
		res[val] = true
	}
	return res
}

// RUN INSTRUCTIONS (RemoveChar):
// ===========================================
// str := "abcdabcdabcd"
// seq := "a"
// fmt.Printf("Original: %s Result:%s\n", str, misc.RemoveChar(str, seq))
//
// str = "abbcdabbcdabbcd"
// seq = "bbc"
// fmt.Printf("Original: %s Result:%s\n", str, misc.RemoveChar(str, seq))
//
//}
// Output
// ===========================================
// Original: abcdabcdabcd Result:bcdbcdbcd
// Original: abbcdabbcdabbcd Result:adadad
