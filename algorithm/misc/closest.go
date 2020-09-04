package misc

import "fmt"

//Closest is based on the below problem
//Given non-zero two integers N and M. The problem is to find the number
//closest to N and divisible by M. If there are more than one such number,
//then output the one having maximum absolute value.
func Closest(n, m int) int {
	var d1, d2, n1, n2 int
	for i := n; i >= n-m; i-- {
		if i%m == 0 {
			n1 = i
			d1 = i - n
			break
		}
	}

	for i := n; i <= n+m; i++ {
		if i%m == 0 {
			n2 = i
			d2 = i - n
			break
		}
	}

	if d1 == d2 {
		fmt.Println(d1, d2, n1, n2)
		if n1 > n2 {
			return n1
		}
		return n2

	} else if d1 < d2 {
		return n1
	} else if d2 < d1 {
		return d2
	}

	return 0
}
