package toh

import "fmt"

func TowerOfHanoi(n int, from, to, aux string) {
	if n == 1 {
		fmt.Printf("Move disk 1 from %s to %s\n", from, to)

		return
	}
	TowerOfHanoi(n-1, from, aux, to)
	fmt.Printf("Move disk %d from %s to %s\n", n, from, to)
	TowerOfHanoi(n-1, aux, to, from)
}
