package interpolation

// =====================================================================================================================
// Interpolation Search
// =====================================================================================================================
// Linear Search finds the element in O(n) time, Jump Search takes O(√ n) time and Binary Search take O(Log n) time.
// The Interpolation Search is an improvement over Binary Search for instances,
// where the values in a sorted array are uniformly distributed.
// Binary Search always goes to the middle element to check.
// On the other hand, interpolation search may go to different locations according
// to the value of the key being searched.
// For example, if the value of the key is closer to the last element,
// interpolation search is likely to start search toward the end side.
//
// To find the position to be searched, it uses following formula:
// ===========================================
// The idea of formula is to return higher value of pos
// when element to be searched is closer to arr[hi]. And
// smaller value when closer to arr[lo]
//
// pos = lo + ((hi - lo) / (arr[hi] - arr[lo])) * (x - arr[lo])
//
// arr[] ==> Array where elements need to be searched
// x     ==> Element to be searched
// lo    ==> Starting index in arr[]
// hi    ==> Ending index in arr[]
//
// Algorithm
// ===========================================
// Rest of the Interpolation algorithm is the same except the above partition logic.
//
// Step1: In a loop, calculate the value of “pos” using the probe position formula.
// Step2: If it is a match, return the index of the item, and exit.
// Step3: If the item is less than arr[pos], calculate the probe position of the left sub-array.
//        Otherwise calculate the same in the right sub-array.
// Step4: Repeat until a match is found or the sub-array reduces to zero.
//
// Time Complexity
// ===========================================
// Time Complexity: If elements are uniformly distributed, then O (log log n)). In worst case it can take upto O(n).
// Auxiliary Space: O(1)
// The worst case complexity of interpolation search is O(n)
// e.g., searching for 1000 in 1,2,...,999,1000,10^9 will take 1000 accesses.
// However, its average case complexity,
// under the assumption that the keys are uniformly distributed, is O(log log n).
//
// When to use
// ===========================================
// Interpolation Search is an improvement over Binary Search for scenarios where
// the values in a sorted array are uniformly distributed.
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// arr := []int{2, 5, 8, 12, 16, 23, 38, 56, 72, 91, 95, 96, 97, 98, 99, 100, 103, 105, 107}
// x := 2
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 5
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 8
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 12
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 16
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 23
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 38
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 56
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 72
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 91
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 95
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 96
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 97
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 98
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 99
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 100
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
//
// x = 22
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
//
// x = 3
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
//
// x = 101
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 102
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 103
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 105
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 107
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
// x = 108
// fmt.Printf("%d is at position %d\n", x, interpolation.Search(arr, x))
//
//}
//
// Output
// ===========================================
// 2 is at position 0
// 5 is at position 1
// 8 is at position 2
// 12 is at position 3
// 16 is at position 4
// 23 is at position 5
// 38 is at position 6
// 56 is at position 7
// 72 is at position 8
// 91 is at position 9
// 95 is at position 10
// 96 is at position 11
// 97 is at position 12
// 98 is at position 13
// 99 is at position 14
// 100 is at position 15
// 22 is at position -1
// 3 is at position -1
// 101 is at position -1
// 101 is at position -1
// 102 is at position -1
// 103 is at position 16
// 105 is at position 17
// 107 is at position 18
// 108 is at position -1
