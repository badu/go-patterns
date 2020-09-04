package docs

// =====================================================================================================================
// In-Place Algorithm
// =====================================================================================================================
// In-place has more than one definitions. One strict definition is.
// An in-place algorithm is an algorithm that does not need an extra space and
// produces an output in the same memory that contains the data by transforming the input ‘in-place’.
// However, a small constant extra space used for variables is allowed.
//
// A more broad definition is,
// In-place means that the algorithm does not use extra space for
// manipulating the input but may require a small though non-constant extra space for its operation.
// Usually, this space is O(log n), though sometimes anything in o(n) (Smaller than linear) is allowed
//
// Which Sorting Algorithms are In-Place and which are not?
// ===========================================
// In Place : Bubble sort, Selection Sort, Insertion Sort, Heapsort.
// Not In-Place : Merge Sort. Note that merge sort requires O(n) extra space.
//
// What about QuickSort? Why is it called In-Place?
// ===========================================
// QuickSort uses extra space for recursive function calls.
// It is called in-place according to broad definition as extra space required is not used to manipulate input,
// but only for recursive calls.
