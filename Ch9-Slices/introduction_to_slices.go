package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)

	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	fmt.Println(nums[3])
	// fmt.Println(nums[5]) panic : this will throw index out of range error
}
