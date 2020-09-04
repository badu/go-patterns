package bubble

// =====================================================================================================================
// Bubble Sort
// =====================================================================================================================
// https://youtu.be/Jdtq5uKz-w4
// Bubble Sort is the simplest sorting algorithm that works by
// repeatedly swapping the adjacent elements if they are in wrong order.
// In each iteration, the largest element is moved or BUBBLED to its rightmost(sorted position) in the array
// So in the next iteration there is no need to compare those elements(sorted elements)
// If the given array has to be sorted in ascending order,
// then bubble sort will start by comparing the first element of the array with the second element,
// if the first element is greater than the second element,
// it will swap both the elements, and then move on to compare the second and the third element, and so on.
// If we have total n elements, then we need to repeat this process for n-1 times.
// It is known as bubble sort, because with every complete iteration the largest element in the given array,
// bubbles up towards the last place or the highest index, just like a water bubble rises up to the water surface.
// Sorting takes place by stepping through all the elements one-by-one and comparing it with
// the adjacent element and swapping them if required.
//
// Example
// ===========================================
// First Pass:
// ( 5 1 4 2 8 ) > ( 1 5 4 2 8 ), Here, algorithm compares the first two elements, and swaps since 5 > 1.
// ( 1 5 4 2 8 ) >  ( 1 4 5 2 8 ), Swap since 5 > 4
// ( 1 4 5 2 8 ) >  ( 1 4 2 5 8 ), Swap since 5 > 2
// ( 1 4 2 5 8 ) > ( 1 4 2 5 8 ), Now, since these elements are already in order (8 > 5), algorithm does not swap them.
//
// Second Pass:
// ( 1 4 2 5 8 ) > ( 1 4 2 5 8 )
// ( 1 4 2 5 8 ) > ( 1 2 4 5 8 ), Swap since 4 > 2
// ( 1 2 4 5 8 ) > ( 1 2 4 5 8 )
// ( 1 2 4 5 8 ) >  ( 1 2 4 5 8 )
// Now, the array is already sorted, but our algorithm does not know if it is completed. The algorithm needs one whole pass without any swap to know it is sorted.
//
// Third Pass:
// ( 1 2 4 5 8 ) > ( 1 2 4 5 8 )
// ( 1 2 4 5 8 ) > ( 1 2 4 5 8 )
// ( 1 2 4 5 8 ) > ( 1 2 4 5 8 )
// ( 1 2 4 5 8 ) > ( 1 2 4 5 8 )
//
// Time Complexity
// ===========================================
// Worst Case Time Complexity [ Big-O ]: O(n^2)
// Best Case Time Complexity [Big-omega]: O(n)
// Average Time Complexity [Big-theta]: O(n^2)
// Auxiliary Space: O(1)
//
// Best case occurs when array is already sorted.
// Worst case occurs when array is reverse sorted.
//
// Boundary Cases: Bubble sort takes minimum time (Order of n) when elements are already sorted.
//
// Sorting In Place: Yes
//
// Stable: Yes
//
// RUN INSTRUCTIONS:
// ===========================================
// func main() {
// arr := []int{95, 80, 30, 90, 50, 40, 70, 10}
// bubble.Sort(arr)
// fmt.Println("Result: ", arr)
//
//}
//
// Output
// ===========================================
// Result:  [10 30 40 50 70 80 90 95]
