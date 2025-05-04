package main

import "fmt"

func main() {
	slice := make([]int, 5)
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
}
