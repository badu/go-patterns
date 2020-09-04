package multiply

// =====================================================================================================================
// Multiply two integers without using multiplication, division and bitwise operators, and no loops
// =====================================================================================================================
// By making use of recursion, we can multiply two integers with the given constraints.
//
// To multiply x and y, recursively add x y times.
//
//
// RUN INSTRUCTIONS :
// ===========================================
// func main() {
// x, y := 10, 20
// fmt.Printf("%d*%d = %d\n", x, y, multiply.NoOperator(x, y))
//
// x, y = 3, -2
// fmt.Printf("%d*%d = %d\n", x, y, multiply.NoOperator(x, y))
//
// x, y = 0, -2
// fmt.Printf("%d*%d = %d\n", x, y, multiply.NoOperator(x, y))
//
//}
// Output
// ===========================================
// 10*20 = 200
// 3*-2 = -6
// 0*-2 = 0
