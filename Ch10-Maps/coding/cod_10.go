/*
10. Group by Remainder
Write a function groupByRemainder(nums []int, divisor int) map[int][]int that groups numbers by their remainder when divided by divisor.
*/
package main

import "fmt"

func groupByRemainder(nums []int, divisor int) map[int][]int {
	res := make(map[int][]int)

	for _, x := range nums {
		val := x % divisor
		res[val] = append(res[val], x)

	}

	return res

}

func main() {
	nums := []int{10, 15, 20, 25, 30, 7, 8, 9}
	divisor := 4

	fmt.Println(groupByRemainder(nums, divisor))
}
