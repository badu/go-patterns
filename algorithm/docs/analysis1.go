package docs

// =====================================================================================================================
// Analysis of Algorithms
// =====================================================================================================================
//
// Asymptotic Analysis
// ===========================================
// Given two algorithms for a task, how do we find out which one is better?
// One naive way of doing this is implement both the algorithms and run the two programs on your computer for
// different inputs and see which one takes less time.
// There are many problems with this approach for analysis of algorithms.
// 1. It might be possible that for some inputs, first algorithm performs better than the second.
//    And for some inputs second performs better
// 2. It might also be possible that for some inputs, first algorithm perform better on one machine and
//    the second works better on other machine for some other inputs
// Asymptotic Analysis is the big idea that handles above issues in analyzing algorithms.
// In Asymptotic Analysis, we evaluate the performance of an algorithm in terms of
// input size (we don’t measure the actual running time). We calculate, how the time (or space)
// taken by an algorithm increases with the input size.
// For example, let us consider the search problem (searching a given item) in a sorted array.
// One way to search is Linear Search (order of growth is linear) and
// the other way is Binary Search (order of growth is logarithmic).
// To understand how Asymptotic Analysis solves the above mentioned problems in analyzing algorithms,
// let us say we run the Linear Search on a fast computer A and Binary Search on a slow computer B and
// we pick the constant values for the two computers so that it tells us exactly how long it
// takes for the given machine to perform the search in seconds. Let’s say the constant for A is 0.2 and
// the constant for B is 1000 which means that A is 5000 times more powerful than B.
// For small values of input array size n, the fast computer may take less time.
// But, after a certain value of input array size,
// the Binary Search will definitely start taking less time compared to the
// Linear Search, even though the Binary Search is being run on a slow machine.
// The reason is the order of growth of Binary Search with respect to input size is
// logarithmic while the order of growth of Linear Search is linear.
// So the machine dependent constants can always be ignored after a certain value of input size.
// Here are some running times for this example:
// Linear Search running time in seconds on A: 0.2 * n
// Binary Search running time in seconds on B: 1000*log(n)
//
//------------------------------------------------
//|n      | Running time on A | Running time on B |
//-------------------------------------------------
//|10     | 2 sec             | ~ 1 h             |
//-------------------------------------------------
//|100    | 20 sec            | ~ 1.8 h           |
//-------------------------------------------------
//|10^6   | ~ 55.5 h          | ~ 5.5 h           |
//-------------------------------------------------
//|10^9   | ~ 6.3 years       | ~ 8.3 h           |
//-------------------------------------------------
//
// Does Asymptotic Analysis always work?
// ===========================================
// Asymptotic Analysis is not perfect, but that’s the best way available for analyzing algorithms.
// For example, say there are two sorting algorithms that take 1000nLogn and 2nLogn time respectively on a machine.
// Both of these algorithms are asymptotically same (order of growth is nLogn).
// So, With Asymptotic Analysis, we can’t judge which one is better as we ignore constants in Asymptotic Analysis.
// Also, in Asymptotic analysis, we always talk about input sizes larger than a constant value.
// It might be possible that those large inputs are never given to your software and
// an algorithm which is asymptotically slower, always performs better for your particular situation.
// So, you may end up choosing an algorithm that is Asymptotically slower but faster for your software.
//
// 1 < log n < sqrt(n) < n < n log n < n^2 < n^3< .........< 2^n < 3^n <....n^n
