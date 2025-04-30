package main

import "fmt"

/*
fmt.Errorf() is like fmt.Sprintf() — but returns an error instead of a string.

It’s used to create dynamic, context-rich error messages.
*/
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by %d", a, b)

	}
	return a / b, nil
}

func main() {
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}
}
