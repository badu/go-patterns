package docs

// =====================================================================================================================
// Divide and Conquer
// =====================================================================================================================
// If we can break a single big problem into smaller sub-problems,
// solve the smaller sub-problems and combine their solutions to find the solution for the original big problem,
// it becomes easier to solve the whole problem.
// Example,
// In Merge Sort, the given unsorted array with n elements, is divided into n subarrays, each having one element,
// because a single element is always sorted in itself. Then, it repeatedly merges these subarrays,
// to produce new sorted subarrays, and in the end, one complete sorted array is produced.
//
// The concept of Divide and Conquer involves three steps:
// 1. Divide the problem into multiple small problems.
// 2. Conquer the subproblems by solving them. The idea is to break down the problem into atomic subproblems,
//    where they are actually solved.
// 3. Combine the solutions of the subproblems to find the solution of the actual problem.
//
// Refer image3
