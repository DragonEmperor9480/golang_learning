/*

Write a function findDuplicates(nums []int) []int that returns a new slice of numbers that appear more than once.

*/

package main

import "fmt"

func findDuplicates(nums []int) []int {
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	dup := []int{}

	for i, k := range freq {
		if k > 1 {
			dup = append(dup, i)
		}
	}

	return dup
}

func main() {
	x := []int{1, 2, 3, 5, 4, 6, 7, 8, 9, 0, 2, 1, 4, 5}
	fmt.Println(findDuplicates(x))
}
