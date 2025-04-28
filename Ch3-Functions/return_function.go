package main

import "fmt"

func mul(a int, b int) int {
	return a * b
}

func main() {
	res := mul(50, 50)
	fmt.Println("Result:", res)
}
