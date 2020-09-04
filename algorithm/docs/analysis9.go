package docs

// =====================================================================================================================
// Analysis of Algorithms-9(Space Complexity)
// =====================================================================================================================
//
//
// What does ‘Space Complexity’ mean?
// ===========================================
// The term Space Complexity is misused for Auxiliary Space at many places.
// Following are the correct definitions of Auxiliary Space and Space Complexity.
// Auxiliary Space is the extra space or temporary space used by an algorithm.
// Space Complexity of an algorithm is total space taken by the algorithm with respect to the input size.
// Space complexity includes both Auxiliary space and space used by input.
// The space complexity of an algorithm or a computer program is the amount of memory space required to solve an
// instance of the computational problem as a function of characteristics of the input.
// It is the memory required by an algorithm to execute a program and produce output.
//
// For example, if we want to compare standard sorting algorithms on the basis of space,
// then Auxiliary Space would be a better criteria than Space Complexity.
// Merge Sort uses O(n) auxiliary space, Insertion sort and Heap Sort use O(1) auxiliary space.
// Space complexity of all these sorting algorithms is O(n) though.
