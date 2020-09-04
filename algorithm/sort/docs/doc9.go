package docs

// =====================================================================================================================
// Sorting Terminology
// =====================================================================================================================
//
// What is in-place sorting?
// ===========================================
// An in-place sorting algorithm uses constant extra space for producing the output (modifies the given array only).
// It sorts the list only by modifying the order of the elements within the list.
// For example,
// Insertion Sort and Selection Sorts are in-place sorting algorithms as they do not use any
// additional space for sorting the list and a typical implementation of Merge Sort is not in-place,
// also the implementation for counting sort is not in-place sorting algorithm.
//
// What is Internal and External Sorting?
// ===========================================
// When all data that needs to be sorted cannot be placed in-memory at a time,
// the sorting is called external sorting. External Sorting is used for massive amount of data.
// Merge Sort and its variations are typically used for external sorting.
// Some external storage like hard-disk, CD, etc is used for external storage.
// When all data is placed in-memory, then sorting is called internal sorting.
//
// External sorting is a term for a class of sorting algorithms that can handle massive amounts of data.
// External sorting is required when the data being sorted do not fit into
// the main memory of a computing device (usually RAM) and instead,
// they must reside in the slower external memory (usually a hard drive).
// External sorting typically uses a hybrid sort-merge strategy.
// In the sorting phase, chunks of data small enough to fit in main memory are
// read, sorted, and written out to a temporary file.
// In the merge phase, the sorted sub-files are combined into a single larger file.
//
// One example of external sorting is the external merge sort algorithm,
// which sorts chunks that each fit in RAM, then merges the sorted chunks together.
// We first divide the file into runs such that the size of a run is small enough to fit into main memory.
// Then sort each run in main memory using merge sort sorting algorithm.
// Finally merge the resulting runs together into successively bigger runs, until the file is sorted.
