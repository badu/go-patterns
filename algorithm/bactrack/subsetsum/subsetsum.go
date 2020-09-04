package subsetsum

import "fmt"

func FindSubset(set []int, size, sum int) {
	subSet := make([]int, size, size) //create subset array to pass parameter of subsetSum
	subsetSum(set, subSet, size, 0, 0, 0, sum)
}

func subsetSum(set, subSet []int, n, subSize, total, nodeCount, sum int) {
	if total == sum {
		//print the subset
		displaySubset(subSet, subSize)
		//for other subsets
		subsetSum(set, subSet, n, subSize-1, total-set[nodeCount-1], nodeCount+1, sum)
		return
	} else {
		for i := nodeCount; i < n; i++ { //find node along breadth
			subSet[subSize] = set[i]
			subsetSum(set, subSet, n, subSize+1, total+set[i], i+1, sum) //do for next node in depth
		}
	}
}

func displaySubset(subSet []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(subSet[i], "  ")
	}
	fmt.Println()
}
