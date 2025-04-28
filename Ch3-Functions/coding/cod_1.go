/*
Write a function multiply that takes 3 integers and returns their product.
*/
package main

import "fmt"

func multiply(a, b, c int) int {
	return a * b * c
}

func main() {
	final_val := multiply(5, 6, 7)
	fmt.Println(final_val)
}
