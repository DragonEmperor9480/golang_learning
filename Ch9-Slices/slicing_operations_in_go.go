package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(nums[1:4])

	//🔹 Full Slice Expression (Go 1.2+)

	s := nums[1:3:8]
	fmt.Println(s)
	fmt.Println("length of s:", len(s))
	fmt.Println("capacity of s:", cap(s))

	//🔹 Slicing Doesn't Copy

	x1 := []int{1, 2, 3, 4, 5, 6}
	x2 := x1[1:3]
	x2[0] = 100
	fmt.Println("x1:", x1)
	fmt.Println("x2:", x2)

	//slicing isnt copying as it modifies the orginal slice

}
