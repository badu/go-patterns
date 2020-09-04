package dynamicprogramming

// =====================================================================================================================
// Dynamic Programming
// =====================================================================================================================
// Dynamic programming is useful for solving optimization problems
// It is basically a way of making your algorithms more efficient bu storing some of the intermediate results
// Dynamic programming says a problem should be solved in sequence of decisions
// Dynamic programming says you should try all possible solutions and pick up the best one
// If a problem has 2^n solutions, DP tries all 2^n solutions, but indirectly!
// Both greedy and dynamic programming are used to solve optimization problems
// The strategy used by both is different but purpose is same
// Optimization problems are those which requires either minimum result or maximum result
// In DP we try to find out all possible solutions and then pick up the best solution
// DP is little time consuming as compared to greedy approach, since we try out all possible approaches
// before arriving at the result
// Mostly DP problem are solved using RECURSIVE FORMULAS
// Though we dont use recursion in programming, but the formulas are recursive
// We can use recursion also, but we mostly use iteration
// DP follows principle of OPTIMALITY
// Principle of OPTIMALITY says that
// a problem can be solved by taking a sequence of decisions to get the optimal solution
// In greedy approach decision is taken one time and then we solve the problem,
// but in DP decision is taken at every stage
// DP Types:
// 1. Tabulation
// 2. Memoization
//
// Dynamic Programming is an algorithmic paradigm that solves a given complex problem by breaking it into
// subproblems and stores the results of subproblems to avoid computing the same results again.
// Following are the two main properties of a problem that suggests that
// the given problem can be solved using Dynamic programming.
// 1) Overlapping Subproblems
// 2) Optimal Substructure
//
// 1. Overlapping Subproblems:
// ===========================================
// Like Divide and Conquer, Dynamic Programming combines solutions to sub-problems.
// Dynamic Programming is mainly used when solutions of same subproblems are needed again and again.
// In dynamic programming, computed solutions to subproblems are stored in a table so that
// these don’t have to be recomputed. So Dynamic Programming is not useful when
// there are no common (overlapping) subproblems because
// there is no point storing the solutions if they are not needed again.
// For example, Binary Search doesn’t have common subproblems.
// If we take an example of following recursive program for Fibonacci Numbers,
// there are many subproblems which are solved again and again.
//
// Simple recursive program for Fibonacci numbers
// int fib(int n)
//{
//   if ( n <= 1 )
//      return n;
//   return fib(n-1) + fib(n-2);
//}
// Recursion tree for execution of fib(5)
//                         fib(5)
//                     /             \
//               fib(4)                fib(3)
//             /      \                /     \
//         fib(3)      fib(2)         fib(2)    fib(1)
//        /     \        /    \       /    \
//  fib(2)   fib(1)  fib(1) fib(0) fib(1) fib(0)
//  /    \
//fib(1) fib(0)
// We can see that the function fib(3) is being called 2 times.
// If we would have stored the value of fib(3), then instead of computing it again,
// we could have reused the old stored value.
//
// There are following two different ways to store the values so that these values can be reused:
// a) Memoization (Top Down)
// b) Tabulation (Bottom Up)
//
// a) Memoization (Top Down)
// ===========================================
// The memoized program for a problem is similar to the recursive version with a small modification that it looks into
// a lookup table before computing solutions. We initialize a lookup array with all initial values as NIL.
// Whenever we need the solution to a subproblem, we first look into the lookup table.
// If the precomputed value is there then we return that value,
// otherwise, we calculate the value and put the result in the lookup table so that it can be reused later.
//
// Following is the memoized version for nth Fibonacci Number:
//
//var lookup = make([]int, 100, 100)
//func fib(n int) int {
//	val := lookup[n]
//	if val == 0 {
//		if n <= 1 {
//			lookup[n] = n
//		} else {
//			lookup[n] = fib(n-1) + fib(n-2)
//		}
//	}
//	return lookup[n]
//}
// Time complexity is O(n+1) = O(n)
//
// Tabulation (Bottom Up)
// ===========================================
// The tabulated program for a given problem builds a table in bottom up fashion and returns the last entry from table.
// For example, for the same Fibonacci number, we first calculate fib(0) then fib(1) then fib(2) then fib(3) and so on.
// So literally, we are building the solutions of subproblems bottom-up.
//
// Following is the tabulated version for nth Fibonacci Number.
// func fib(n int) int {
//
//	f := make([]int, n+1, n+1)
//	f[0] = 0
//	f[1] = 1
//	for i := 2; i <= n; i++ {
//		f[i] = f[i-1] + f[i-2]
//	}
//	return f[n]
//}
//
// Both Tabulated and Memoized store the solutions of subproblems.
// In Memoized version, table is filled on demand while in Tabulated version,
// starting from the first entry, all entries are filled one by one. Unlike the Tabulated version,
// all entries of the lookup table are not necessarily filled in Memoized version.
// For example, Memoized solution of the LCS problem doesn’t necessarily fill all entries.
// Time taken by Recursion method is much more than the two
// Dynamic Programming techniques mentioned above  Memoization and Tabulation!
//
// Optimal Substructure
// ===========================================
// A given problems has Optimal Substructure Property if
// optimal solution of the given problem can be obtained by using optimal solutions of its subproblems.
// For example, the Shortest Path problem has following optimal substructure property:
// If a node x lies in the shortest path from a source node u to destination node v then
// the shortest path from u to v is combination of shortest path from u to x and shortest path from x to v.
// The standard All Pair Shortest Path algorithms like
// FloydWarshall and BellmanFord are typical examples of Dynamic Programming.
// On the other hand, the Longest Path problem doesn’t have the Optimal Substructure property.
// Here by Longest Path we mean longest simple path (path without cycle) between two nodes.
// Consider the following unweighted graph given in the CLRS book.
// There are two longest paths from q to t: q→r→t and q→s→t.
// Unlike shortest paths, these longest paths do not have the optimal substructure property.
// For example, the longest path q→r→t is not a combination of longest path from q to r and longest path from r to t,
// because the longest path from q to r is q→s→t→r and the longest path from r to t is r→q→s→t.
// Refer image1
