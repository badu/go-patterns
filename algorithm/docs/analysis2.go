package docs

// =====================================================================================================================
// Analysis of Algorithms-2
// =====================================================================================================================
//
// Worst, Average and Best Cases
// ===========================================
// In this post, we will take an example of Linear Search and analyze it using Asymptotic analysis.
// We can have three cases to analyze an algorithm:
// 1. Worst Case
// 2. Average Case
// 3. Best Case
//
// Let us consider the following implementation of Linear Search:
//
// func search(arr []int, val int) int {
//	for pos, v := range arr {
//		if v == val {
//			return pos + 1
//		}
//	}
//	return -1
// }
// func main() {
//	arr := []int{1, 10, 30, 15}
//	val := 30
//	fmt.Printf("%d is at position %d\n", val, search(arr, val))
// }
//
//
// Worst Case Analysis (Usually Done)
// ===========================================
// In the worst case analysis, we calculate upper bound on running time of an algorithm.
// We must know the case that causes maximum number of operations to be executed.
// For Linear Search, the worst case happens when the element to be searched (val in the above code)
// is not present in the array. When x is not present,
// the search() functions compares it with all the elements of arr[] one by one.
// Therefore, the worst case time complexity of linear search would be Θ(n).
//
// Average Case Analysis (Sometimes done)
// ===========================================
// In average case analysis, we take all possible inputs and calculate computing time for all of the inputs.
// Sum all the calculated values and divide the sum by total number of inputs.
// We must know (or predict) distribution of cases.
// For the linear search problem, let us assume that all cases are uniformly distributed
// (including the case of x not being present in array).
// So we sum all the cases and divide the sum by (n+1).
// Following is the value of average case time complexity.
// (Refer image 1)
//
// Best Case Analysis (Bogus)
// ===========================================
// In the best case analysis, we calculate lower bound on running time of an algorithm.
// We must know the case that causes minimum number of operations to be executed.
// In the linear search problem, the best case occurs when x is present at the first location.
// The number of operations in the best case is constant (not dependent on n).
// So time complexity in the best case would be Θ(1)
//
// Summary
// ===========================================
// Most of the times, we do worst case analysis to analyze algorithms.
// In the worst analysis, we guarantee an upper bound on the running time of an algorithm which is good information.
// The average case analysis is not easy to do in most of the practical cases and it is rarely done.
// In the average case analysis, we must know (or predict) the mathematical distribution of all possible inputs.
// The Best Case analysis is bogus. Guaranteeing a lower bound on
// an algorithm doesn’t provide any information as in the worst case, an algorithm may take years to run.
//
// For some algorithms, all the cases are asymptotically same, i.e., there are no worst and best cases.
// For example, Merge Sort. Merge Sort does Θ(nLogn) operations in all cases.
// Most of the other sorting algorithms have worst and best cases.
// For example, in the typical implementation of Quick Sort (where pivot is chosen as a corner element),
// the worst occurs when the input array is already sorted and the best occur when the pivot elements always
// divide array in two halves.
// For insertion sort, the worst case occurs when the array is reverse sorted and
// the best case occurs when the array is sorted in the same order as output.
