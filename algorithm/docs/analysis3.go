package docs

// =====================================================================================================================
// Analysis of Algorithms-3
// =====================================================================================================================
//
// Asymptotic Notations
// ===========================================
// The main idea of asymptotic analysis is to have a measure of efficiency of algorithms that doesn’t
// depend on machine specific constants, and doesn’t require algorithms to be implemented and time
// taken by programs to be compared.
// Asymptotic notations are mathematical tools to represent time complexity of algorithms for asymptotic analysis.
//
// The following 3 asymptotic notations are mostly used to represent time complexity of algorithms.
//
// Θ Notation (refer image2)
// ===========================================
// The theta notation bounds a functions from above and below, so it defines exact asymptotic behavior.
// A simple way to get Theta notation of an expression is to drop low order terms and ignore leading constants.
// For example, consider the following expression:
// 3n3 + 6n2 + 6000 = Θ(n3)
// Dropping lower order terms is always fine because there will always be a n0 after which Θ(n^3) has
// higher values than Θ(n^2) irrespective of the constants involved.
// For a given function g(n), we denote Θ(g(n)) is following set of functions: (image3)
// The (image3) definition means, if f(n) is theta of g(n),
// then the value f(n) is always between c1*g(n) and c2*g(n) for large values of n (n >= n0).
// The definition of theta also requires that f(n) must be non-negative for values of n greater than n0.
//
// Big O Notation (refer image4)
// ===========================================
// Big O notation is used to measure how running time or space requirements for your program grows as
// input size grows
// The Big O notation defines an upper bound of an algorithm, it bounds a function only from above.
// For example, consider the case of Insertion Sort.
// It takes linear time in best case and quadratic time in worst case.
// We can safely say that the time complexity of Insertion sort is O(n^2).
// Note that O(n^2) also covers linear time.
// If we use Θ notation to represent time complexity of Insertion sort,
// we have to use two statements for best and worst cases:
// 1. The worst case time complexity of Insertion Sort is Θ(n^2).
// 2. The best case time complexity of Insertion Sort is Θ(n).
// The Big O notation is useful when we only have upper bound on time complexity of an algorithm.
// Many times we easily find an upper bound by simply looking at the algorithm:(refer image5)
// To get the big O notation:
// 1. We consider the fastest growing term in the equation and discard the other terms.
// 2. From the fastest growing term obtained from step 1 we discard(or drop) the constant(if any)
// 3. The leftover part can now be represented as the big O notation
// Example:
// if 2*a^2 + 10* b + c1 + c2 is the function that represents our algorithm,
// After applying step 1 we get: 2*a^2
// After step 2: a^2
// Thus the big O notation of our algorithm is: O(a^2) or O(n^2) by convention
//
// The reason we drop constants in big O notation is because, in big O notation we the focus is on the BEHAVIOUR
// not the actual value. And for large values of input the factor of constants become negligible(but the behaviour of
// the graph will be same)
//
//
// Ω Notation (refer image6)
// ===========================================
// Just as Big O notation provides an asymptotic upper bound on a function,
// Ω notation provides an asymptotic lower bound.
// Ω Notation can be useful when we have lower bound on time complexity of an algorithm.
// As discussed in the previous post, the best case performance of an algorithm is generally not useful,
// the Omega notation is the least used notation among all three.
// For a given function g(n), we denote by Ω(g(n)) the set of functions: (refer image7)
// Let us consider the same Insertion sort example here.
// The time complexity of Insertion Sort can be written as Ω(n),
// but it is not a very useful information about insertion sort,
// as we are generally interested in worst case and sometimes in average case.
//
//
// Best, Average and worst cases are not related to the notations(asymptotic notations), notations are for functions,
// (mathematical functions) cases are the type of analysis done on algorithms
// So either of the Big O, Omega or Theta can be used to represent best, average and worst cases.
//
// Whenever you have a function if it possible to represent it with a Theta it would be better
// If it is not possible we can go for Big O or Omega.
