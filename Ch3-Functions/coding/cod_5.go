/*

Write a recursive function that calculates the sum of numbers from 1 to n.

*/

package main

import "fmt"

func RecSum(n int) int {
	if n == 1 {
		return 1
	}
	return n + RecSum(n-1)

}

func main() {
	fmt.Println(RecSum(5))

}
