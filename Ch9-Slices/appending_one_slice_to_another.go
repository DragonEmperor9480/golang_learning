package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 0}
	a = append(a, b...) // use ... to unpack elements of b
	fmt.Println(a)      // [1 2 3 4 5]
}
