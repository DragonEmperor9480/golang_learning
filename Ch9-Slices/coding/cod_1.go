/*

Write a function removeEven(nums []int) []int that returns a new slice containing only the odd numbers from the input.

*/

package main

import "fmt"

func removeEven(nums []int) []int {
	x := []int{}
	for i := 0; i < len(nums); i++ {
		if (nums[i] % 2) != 0 {
			x = append(x, nums[i])
		}

	}
	return x
}

func main() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(removeEven(s1))

}
