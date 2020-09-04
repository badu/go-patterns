package epower

// =====================================================================================================================
// Efficient program to calculate e^x
// =====================================================================================================================
// The value of Exponential Function e^x can be expressed using following Taylor Series.
// e^x = 1 + x + x^2 / 2! + x^3 / 3! + ... + every x^n / n!
//
// How to efficiently calculate the sum of above series?
// The series can be re-written as:
// e^x = 1 + (x/1) (1 + (x/2) (1 + (x/3) (........) ) )
//
// Let the sum needs to be calculated for n terms, we can calculate sum using following loop.
// for (i = n - 1; i > 0; --i )
//    sum = 1 + x * sum / i;
//
// RUN INSTRUCTIONS :
// ===========================================
// func main() {
// n := 10
// x := 1.0
// fmt.Printf("e^*%f = %f\n", x, epower.Pow(n, x))
//
//}
// Output
// ===========================================
// e^*1.000000 = 2.718282
