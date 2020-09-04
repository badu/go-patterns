package median

// =====================================================================================================================
// Median of two sorted arrays of SAME size
// =====================================================================================================================
// There are 2 sorted arrays A and B of size n each.
// Write an algorithm to find the median of the array obtained after
// merging the above 2 arrays(i.e. array of length 2n).
// The complexity should be O(log(n)).
// Since size of the set for which we are looking for median is even (2n),
// we need take average of middle two numbers and return floor of the average.
//
// Method 1 (Simply count while Merging)
// ===========================================
// Use merge procedure of merge sort. Keep track of count while comparing elements of two arrays.
// If count becomes n(For 2n elements), we have reached the median.
// Take the average of the elements at indexes n-1 and n in the merged array.
//
// Method 2 (By comparing the medians of two arrays)
// ===========================================
// This method works by first getting medians of the two sorted arrays and then comparing them.
// Let ar1 and ar2 be the input arrays.
//
// Example
// ===========================================
// Input: arr1[] = {1, 12, 15, 26, 38}
//        arr2[] = {2, 13, 17, 30, 45}
// Output: 16
// Explanation:
// After merging two arrays we get: {1, 2, 12, 13, 15, 17, 26, 30, 38, 45}
// Middle two elements are 15, 17
// Average of middle elements is: (15+17)/2=16
//
//
// RUN INSTRUCTIONS :
// ===========================================
// func main() {
//
//
//}
// Output
// ===========================================
//
