/*
🧩 1. Sum of Array Elements
Task:
Write a function sumArray(arr [5]int) int that takes a fixed-size array of 5 integers and returns the sum of its elements.
*/
package main

import "fmt"

func sumArray(arr [5]int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func main() {
	numbers := [5]int{1, 2, 3, 4, 5}
	res := sumArray(numbers)
	fmt.Println(res)

}
