package misc

import (
	"fmt"
)

//PrintTriangle prints triangle for a given level
//levels>0
//    *
//   * *
//  * * *
// * * * *
//* * * * *
//Print stars at,,
//at 5
//at 4,6
//at 3,5,7
//at 2,4,6,8
//at 1,3,5,7,9
func PrintTriangle(levels int) {
	starsToPrint := 1
	//Max maxSteps to be taken in each level
	maxSteps := levels
	for i := 1; i <= levels; i++ {
		tmpMap := generateStarsLocationMap(starsToPrint, levels-i)
		for j := 1; j <= maxSteps; j++ {
			if tmpMap[j] == j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		maxSteps++
		starsToPrint++
		fmt.Println()
	}
}

//Generates a map with indexes at which the stars have to be printed at each level
func generateStarsLocationMap(count int, lTmp int) map[int]int {
	tmpMap := make(map[int]int)

	for i := 0; i < count; i++ {
		tmpMap[lTmp] = lTmp
		lTmp = lTmp + 2
	}
	return tmpMap
}
