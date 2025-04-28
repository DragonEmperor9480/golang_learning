package main

import "fmt"

func findMax(numbers ...int) int {
	max := numbers[0]
	for _, num := range numbers {
		if max < num {
			max = num
		}
	}
	return max
}

func main() {
	x := findMax(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Max number:", x)
}
