package ones

import "fmt"

func Complement(n byte) {

	// 11111111
	var compByte byte = 255

	fmt.Printf("Input : %08b\n", n)
	fmt.Printf("Output: %08b\n", compByte^n)
}
