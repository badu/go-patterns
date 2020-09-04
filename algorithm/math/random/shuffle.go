package random

import (
	"math/rand"
	"sort"
	"time"
)

func Shuffle(arr []int) {

	sort.Ints(arr)
	tMap := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		tMap[arr[i]] = true
	}
	random(tMap, arr[len(arr)-1], arr)
}

func random(tMap map[int]bool, high int, arr []int) {

	count := 0
	rand.Seed(time.Now().UnixNano())
	for count < len(arr) {
		rNum := rand.Intn(high + 1)
		if tMap[rNum] {
			tMap[rNum] = false
			arr[count] = rNum
			count++
		}
	}
}
