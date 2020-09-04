package prime

// =====================================================================================================================
// Sieve of Eratosthenes for prime numbers
// =====================================================================================================================
// Given a number n, to print all primes smaller than or equal to n.
// Also assuming that n is a small number.
// The sieve of Eratosthenes is one of the most efficient ways to find all primes smaller than n when
// n is smaller than 10 million or so
// Following is the algorithm to find all the prime numbers less than or
// equal to a given integer n by the Eratosthene’s method:
// 1. Create a list of consecutive integers from 2 to n: (2, 3, 4, …, n).
// 2. Initially, let p equal to 2, the first prime number.
// 3. Starting from p^2, count up in increments of p and
//    mark each of these numbers greater than or equal to p^2 itself in the list.
//    These numbers will be p(p+1), p(p+2), p(p+3), etc..
// 4. Find the first number greater than p in the list that is not marked.
//    If there was no such number, stop.
//   Otherwise, let p now equal this number (which is the next prime), and repeat from step 3.
// When the algorithm terminates, all the numbers in the list that are not marked are prime.
//
// Explanation with Example:
// ===========================================
// Let us take an example when n = 50. So we need to print all prime numbers smaller than or equal to 50.
//
// We create a list of all numbers from 2 to 50.
// Refer image1
// According to the algorithm we will mark all the numbers which are divisible by 2 and
// are greater than or equal to the square of it.
// Refer image2
// Now we move to our next unmarked number 3 and mark all the numbers which are multiples of 3 and
// are greater than or equal to the square of it.
// Refer image3
// We move to our next unmarked number 5 and mark all multiples of 5 and
// are greater than or equal to the square of it.
// Refer image4
// We continue this process and our final table will look like below:
// Refer image5
// So the prime numbers are the unmarked ones: 2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47.
// As a extension to this problem, to find the sum of n prime numbers, run a loop from 1 to n
// after applying the Eratosthenes method
//
// Examples
// ===========================================
// Input : n =10
// Output : 2 3 5 7
//
// Input : n = 20
// Output: 2 3 5 7 11 13 17 19
//
// For 23 it’s 2 so 23 is not a multiple of 3
//
// Time Complexity Example
// ===========================================
//
//
// RUN INSTRUCTIONS 1:
// ===========================================
// func main() {
// n := 50
// fmt.Printf("Primes in range %d are %v\n", n, prime.Eratosthenes(n))
//}
// Output
// ===========================================
// Primes in range 50 are [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47]
