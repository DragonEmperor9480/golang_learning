package main

import "fmt"

/*
Named Return Values

Sometimes Go functions can name their return variables in the signature itself.
It can make the code cleaner and no need to explicitly declare variables inside the function.
*/
func areaAndPerimeter(a int, b int) (area int,perimeter int) {
	area = a * b
	perimeter = 2 * (a + b)
	return
}

func main() {
	x, y := areaAndPerimeter(10, 10)
	fmt.Printf("Area: %d\n", x)
	fmt.Printf("Perimeter: %d\n", y)
}
