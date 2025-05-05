/*
Write a function sumSlice(nums []int) int that returns the sum of all elements in a slice.
*/

package main

import "fmt"

func sumSlice(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println("Sum:", sumSlice(s1))
}
