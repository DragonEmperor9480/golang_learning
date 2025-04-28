package main

import "fmt"

//multiple returns

func divideAndRemainder(a int, b int) (int, int) {
	quo := a / b
	rem := a % b
	return quo, rem
}

func main() {
	x, y := divideAndRemainder(10, 5)
	fmt.Println("Quotient:", x)
	fmt.Println("Remainder:", y)
}
