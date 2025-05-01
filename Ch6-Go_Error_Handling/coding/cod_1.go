package main

import (
	"errors"
	"fmt"
)
/*

🧩 1. Division with Error Handling
Task:
Write a function Divide(a, b int) (int, error) that returns an error if b is zero.
In main(), call this function with valid and invalid inputs.

*/

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Not Divisible by zero")
	}
	return a / b, nil
}

func main() {
	res, err := Divide(10, 0)
	if err != nil {
		fmt.Println("Error:", res)
	} else {
		fmt.Println("Result:", res)
	}
}
