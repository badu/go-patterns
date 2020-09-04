package docs

// =====================================================================================================================
// Sort Algorithms Questions-1
// =====================================================================================================================
//
// When does the worst case of QuickSort occur?
// ===========================================
// In quickSort, we select a pivot element,
// then partition the given array around the pivot element by placing
// pivot element at its correct position in sorted array.
// The worst case of quickSort occurs when one part after partition contains all elements and other part is empty.
// For example, if the input array is sorted and if last or first element is chosen as a pivot, then the worst occurs
