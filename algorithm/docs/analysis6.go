package docs

// =====================================================================================================================
// Analysis of Algorithms-6
// =====================================================================================================================
//
// O(1)
// ===========================================
// Time complexity of a function (or set of statements) is considered as O(1) if it doesn’t contain loop,
// recursion and call to any other non-constant time function.
// For example swap() function has O(1) time complexity.
// A loop or recursion that runs a constant number of times is also considered as O(1).
// For example the following loop is O(1).
//
// // Here c is a constant
//   for (int i = 1; i <= c; i++) {
//        // some O(1) expressions
//   }
//
// O(n)
// ===========================================
// Time Complexity of a loop is considered as O(n) if the loop variables is
// incremented / decremented by a constant amount.
// For example following functions have O(n) time complexity.
//    // Here c is a positive integer constant
//
//   for (int i = 1; i <= n; i += c) {
//        // some O(1) expressions
//   }
//
//   for (int i = n; i > 0; i -= c) {
//        // some O(1) expressions
//   }
//
// O(n^c)
// ===========================================
// Time complexity of nested loops is equal to the number of times the innermost statement is executed.
// For example the following sample loops have O(n^2) time complexity
//Ex 1:
// for (int i = 1; i <=n; i += c) {
//       for (int j = 1; j <=n; j += c) {
//          // some O(1) expressions
//       }
//   }
//
// Ex 2:
//   for (int i = n; i > 0; i -= c) {
//       for (int j = i+1; j <=n; j += c) {
//          // some O(1) expressions
//   }
// For example Selection sort and Insertion Sort have O(n^2) time complexity.
//
// O(Log n)
// ===========================================
// Time Complexity of a loop is considered as O(Logn) if the loop variables is
// divided / multiplied by a constant amount.
// Ex 1:
// for (int i = 1; i <=n; i *= c) {
//       // some O(1) expressions
//   }
// Ex 2:
//   for (int i = n; i > 0; i /= c) {
//       // some O(1) expressions
//   }
// For example Binary Search has O(Log n) time complexity.
// Let us see mathematically how it is O(Log n).
// The series that we get in first loop is 1, c, c^2, c3^, … c^k.
// If we put k equals to Logc n, we get c ^ Logc n which is n.
//
// O(Log Log n)
// ===========================================
// Time Complexity of a loop is considered as O(Log Log n) if the loop variables is
// reduced / increased exponentially by a constant amount.
// Ex 1:
// // Here c is a constant greater than 1
//   for (int i = 2; i <=n; i = pow(i, c)) {
//       // some O(1) expressions
//   }
// Ex 2:
//   //Here fun is sqrt or cuberoot or any other constant root
//   for (int i = n; i > 1; i = fun(i)) {
//       // some O(1) expressions
//   }
//
// How to combine time complexities of consecutive loops?
// ===========================================
// When there are consecutive loops, we calculate time complexity as sum of time complexities of individual loops.
// for (int i = 1; i <=m; i += c) {
//        // some O(1) expressions
//   }
//   for (int i = 1; i <=n; i += c) {
//        // some O(1) expressions
//   }
// Time complexity of above code is O(m) + O(n) which is O(m+n)
// If m == n, the time complexity becomes O(2n) which is O(n).
//
// How to calculate time complexity when there are many if, else statements inside loops?
// ===========================================
// As discussed, worst case time complexity is the most useful among best, average and worst.
// Therefore we need to consider worst case.
// We evaluate the situation when values in if-else conditions cause maximum number of statements to be executed.
// For example consider the linear search function where we consider the case when
// element is present at the end or not present at all.
// When the code is too complex to consider all if-else cases,
// we can get an upper bound by ignoring if else and other complex control statements.
//
// How to calculate time complexity of recursive functions?
// ===========================================
// Time complexity of a recursive function can be written as a mathematical recurrence relation.
// To calculate time complexity, we must know how to solve recurrences.
// We will soon be discussing recurrence solving techniques as a separate post.
//
// Difference between Posteriori and Priori analysis
// ===========================================
// Algorithm is a combination or sequence of finite-state to solve a given problem.
// If the problem is having more than one solution or algorithm then the best one is decided by
// the analysis based on two factors.
// 1. CPU Time (Time Complexity)
// 2. Main memory space (Space Complexity)
// Time complexity of an algorithm can be calculated by using two methods:
// 1. Posteriori Analysis
// 2. Priori Analysis
//
// POSTERIORI ANALYSIS
// ===========================================
// Posteriori analysis is a relative analysis.
// It is dependent on language of compiler and type of hardware.
// It will give exact answer.
// It doesn’t use asymptotic notations to represent the time complexity of an algorithm.
// The time complexity of an algorithm using a posteriori analysis differ from system to system.
// If the time taken by the algorithm is less, then the credit will go to compiler and hardware.
//
// PRIORI ANALYSIS
// ===========================================
// Piori analysis is an absolute analysis.
// It is independent of language of compiler and types of hardware.
// It will give approximate answer.
// It uses the asymptotic notations to represent how much time the
// algorithm will take in order to complete its execution.
// The time complexity of an algorithm using a priori analysis is same for every system.
// If the program running faster, credit goes to the programmer.
