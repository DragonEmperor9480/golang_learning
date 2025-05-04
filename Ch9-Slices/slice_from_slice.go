package main

import "fmt"

func main() {
	//slice from slice
	base := []int{10, 20, 30, 40, 50}
	slice1 := base[1:4]
	slice2 := slice1[1:3]

	fmt.Println("Slice 1:", slice1)
	fmt.Println("Slice 2:", slice2)
	fmt.Println("len(slice2):", len(slice2))
	fmt.Println("cap(slice2):", cap(slice2))
}
