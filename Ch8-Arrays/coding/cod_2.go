/*
🧩 2. Find Maximum Element
Task:
Write a function findMax(arr [6]int) int that returns the maximum value in the given array.
*/

package main

import "fmt"

func findMax(arr [6]int) int {
	max := arr[0]
	for _, x := range arr {
		if max < x {
			max = x
		}
	}
	return max
}

func main() {
	numbers := [6]int{10, 20, 30, 40, 60, 100}
	fmt.Println("Max number:", findMax(numbers))
}
