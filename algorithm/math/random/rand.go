package random

import (
	"math/rand"
	"time"
)

//Program to generate random number sequence
//Sequence: min<=Nmin, n2....Nmax <=max
//using built in functions in go
func RandGenerator(min, max int) []int {
	//By default seed is 1
	//If seed is not specified same sequence
	//of rand number will be generated
	rand.Seed(time.Now().UnixNano())

	numGenerated := 0
	res := make([]int, 0, 1)
	tmpMap := make(map[int]bool)
	for numGenerated < max {
		//if max=10 and min=1
		//max-min+1 ensures the number is between 1 & 10
		//with max-min the number will be between 1 & 9
		num := rand.Intn(max-min+1) + min
		if !tmpMap[num] {
			numGenerated++
			tmpMap[num] = true
			res = append(res, num)
		}
	}
	return res
}
