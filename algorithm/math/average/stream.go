package average

import "fmt"

func Stream(arr []int, n int) float64 {
	sum := 0
	avg := 0.0
	for i := 0; i < n; i++ {
		sum = sum + arr[i]
		avg = float64(sum) / float64(i+1)
		fmt.Printf("Average of %d numbers is %f\n", i+1, avg)
	}
	return avg
}
