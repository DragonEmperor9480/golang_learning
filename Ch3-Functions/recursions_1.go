package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(5))
}

// how it works

/*
Step			What Happens
factorial(5)	5 × factorial(4)
factorial(4)	4 × factorial(3)
factorial(3)	3 × factorial(2)
factorial(2)	2 × factorial(1)
factorial(1)	1 × factorial(0)
factorial(0)	returns 1 (base case)

1 → 1×1 → 2×1 → 3×2 → 4×6 → 5×24 = 120

*/
