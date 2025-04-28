package main

import "fmt"

/*
Variadic Functions
Variadic functions can accept an unlimited number of arguments!

*/

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	result := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("sum:", result)
}
