package docs

// =====================================================================================================================
// Analysis of Algorithms-7(Solving Recurrences)
// =====================================================================================================================
// Many algorithms are recursive in nature.
// When we analyze them, we get a recurrence relation for time complexity.
// We get running time on an input of size n as a function of n and the running time on inputs of smaller sizes.
// For example in Merge Sort,
// to sort a given array, we divide it in two halves and recursively repeat the process for the two halves.
// Finally we merge the results.
// Time complexity of Merge Sort can be written as T(n) = 2T(n/2) + cn.
// There are many other algorithms like Binary Search, Tower of Hanoi, etc.
//
// There are mainly three ways for solving recurrences.
//
// 1) Substitution Method
// ===========================================
// We make a guess for the solution and then we use mathematical induction to prove the guess is correct or incorrect.
// For example consider the recurrence T(n) = 2T(n/2) + n
//
// We guess the solution as T(n) = O(nLogn). Now we use induction
// to prove our guess.
//
// We need to prove that T(n) <= cnLogn. We can assume that it is true
// for values smaller than n.
//
// T(n) = 2T(n/2) + n
//    <= 2cn/2Log(n/2) + n
//    =  cnLogn - cnLog2 + n
//    =  cnLogn - cn + n
//    <= cnLogn
//
// 2) Recurrence Tree Method
// ===========================================
// In this method, we draw a recurrence tree and calculate the time taken by every level of tree.
// Finally, we sum the work done at all levels.
// To draw the recurrence tree, we start from the given recurrence and keep drawing till we find a pattern among levels.
// The pattern is typically a arithmetic or geometric series.
// For example consider the recurrence relation
//T(n) = T(n/4) + T(n/2) + cn2
//
//           cn2
//         /      \
//     T(n/4)     T(n/2)
//
//If we further break down the expression T(n/4) and T(n/2),
//we get following recursion tree.
//
//                cn2
//           /           \
//       c(n2)/16      c(n2)/4
//      /      \          /     \
//  T(n/16)     T(n/8)  T(n/8)    T(n/4)
//Breaking down further gives us following
//                 cn2
//            /            \
//       c(n2)/16          c(n2)/4
//       /      \            /      \
//c(n2)/256   c(n2)/64  c(n2)/64    c(n2)/16
// /    \      /    \    /    \       /    \
//
//To know the value of T(n), we need to calculate sum of tree
//nodes level by level. If we sum the above tree level by level,
//we get the following series
//T(n)  = c(n^2 + 5(n^2)/16 + 25(n^2)/256) + ....
//The above series is geometrical progression with ratio 5/16.
//
//To get an upper bound, we can sum the infinite series.
//We get the sum as (n2)/(1 - 5/16) which is O(n2)
//
// 3) Master Method
// ===========================================
// Master Method is a direct way to get the solution.
// The master method works only for following type of recurrences or
// for recurrences that can be transformed to following type.
// T(n) = aT(n/b) + f(n) where a >= 1 and b > 1
// There are following three cases:
// 1. If f(n) = Θ(n^c) where c < Logba then T(n) = Θ(n^Logba)
//
// 2. If f(n) = Θ(n^c) where c = Logba then T(n) = Θ(n^c Log n)
//
// 3.If f(n) = Θ(n^c) where c > Logba then T(n) = Θ(f(n))
//
// How does this work?
// Master method is mainly derived from recurrence tree method.
// If we draw recurrence tree of T(n) = aT(n/b) + f(n),
// we can see that the work done at root is f(n) and work done at all leaves is Θ(nc) where c is Logba.
// And the height of recurrence tree is Logbn
// Refer image9
// In recurrence tree method, we calculate total work done.
// If the work done at leaves is polynomially more, then leaves are the dominant part,
// and our result becomes the work done at leaves (Case 1).
// If work done at leaves and root is asymptotically same,
// then our result becomes height multiplied by work done at any level (Case 2).
// If work done at root is asymptotically more, then our result becomes work done at root (Case 3).
