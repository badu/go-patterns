package duplicate

//RemoveDuplicates removes duplicates from the given array
//Assumes an int array is passed
//Similar logic can be used for other data types as well
func RemoveDuplicates(arr []int) []int {
	isDuplicate := make(map[int]bool)
	res := make([]int, 0, 1)
	for _, val := range arr {
		if !isDuplicate[val] {
			isDuplicate[val] = true
			res = append(res, val)
		}
	}
	return res
}
