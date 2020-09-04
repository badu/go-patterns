package gcd

// =====================================================================================================================
// GCD or HCF of two numbers
// =====================================================================================================================
// GCD (Greatest Common Divisor) or HCF (Highest Common Factor) of two numbers is
// the largest number that divides both of them.
// For example GCD of 20 and 28 is 4 and GCD of 98 and 56 is 14.
//
// A simple solution is to find all prime factors of both numbers,
// then find intersection of all factors present in both numbers.
// Finally return product of elements in the intersection.
//
// An efficient solution is to use Euclidean algorithm which is the main algorithm used for this purpose.
// The idea is, GCD of two numbers doesn’t change if smaller number is subtracted from a bigger number.
//
// The recursive Euclid’s algorithm computes the GCD by using a pair of positive integers
// a and b and returning b and a%b till b is zero.
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// a,b := 98, 56
// fmt.Printf("GCD of %d and %d is %d\n", a, b, gcd.GCD(a,b))
//}
// Output
// ===========================================
// GCD of 98 and 56 is 14
