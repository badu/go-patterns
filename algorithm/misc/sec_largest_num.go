package misc

//Returns the second largest number in the array
//Sorting using in built functions not allowed
//Function expects len of arr > 1
func GetSecondLargestNumInArray(arr []int) int {
	high := 0
	secHigh := 0
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			high = arr[0]
		} else if arr[i] > high {
			secHigh = high
			high = arr[i]
		} else if arr[i] > secHigh {
			secHigh = arr[i]
		}
	}
	return secHigh
}
