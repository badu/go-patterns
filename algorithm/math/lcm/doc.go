package lcm

// =====================================================================================================================
// LCM of two numbers
// =====================================================================================================================
// LCM (Least Common Multiple) of two numbers is the smallest number which can be divided by both numbers.
// For example LCM of 15 and 20 is 60 and LCM of 5 and 7 is 35.
// A simple solution is to find all prime factors of both numbers,
// then find union of all factors present in both numbers.
// Finally return product of elements in union.
// An efficient solution is based on below formula for LCM of two numbers ‘a’ and ‘b’.
// LCM(a, b) = (a x b) / GCD(a, b)
//
// Without using GCD
// Finding LCM using GCD is explained above but you can also find LCM without using GCD
// The approach is to start with the largest of the 2 numbers and
// keep incrementing the larger number by itself till smaller number perfectly divides the resultant.
//
// Examples
// ===========================================
//
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// a, b := 15, 20
// fmt.Printf("LCM of %d and %d is %d\n", a, b, lcm.ByGCD(a, b))
//
//}
// Output
// ===========================================
// LCM of 15 and 20 is 60
