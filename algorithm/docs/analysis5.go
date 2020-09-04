package docs

// =====================================================================================================================
// Analysis of Algorithms-5
// =====================================================================================================================
//
// Lower and Upper Bound Theory
// ===========================================
// The Lower and Upper Bound Theory provides a way to find the lowest complexity algorithm to solve a problem.
// Before understanding the theory, first lets have a brief look on what actually Lower and Upper bounds are.
//
// 1. Lower Bound
//    Let L(n) be the running time of an algorithm A(say), then g(n) is the Lower Bound of A if
//    there exist two constants C and N such that L(n) >= C*g(n) for n > N.
// 2. Upper Bound
//    Let U(n) be the running time of an algorithm A(say), then g(n) is the Upper Bound of A if
//    there exist two constants C and N such that U(n) <= C*g(n) for n > N.
//    Upper bound of an algorithm is shown by the asymptotic notation called Big Oh(O) (or just Oh).
//
// For a more detailed explanation refer:
// https://youtu.be/A03oI0znAoc
// https://youtu.be/Nd0XDY-jVHs
// https://youtu.be/lj3E24nnPjI
