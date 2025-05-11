/*
6. Count Duplicates in Slice
Write a function findDuplicates(nums []int) []int that returns all numbers that appear more than once using a map.
*/
package main

import "fmt"

func findDuplicates(nums []int) []int {
	countingMap := make(map[int]int)
	countingSlice := []int{}
	for _, x := range nums {
		countingMap[x]++

	}
	for key, value := range countingMap {
		if value > 1 {
			countingSlice = append(countingSlice, key)
		}

	}

	return countingSlice
}

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 1, 2, 1, 2, 3}
	fmt.Println(findDuplicates(x))

}
