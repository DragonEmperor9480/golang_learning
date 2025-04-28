package main

import "fmt"

/*
Functions as First-Class Citizens
In Go, functions are first-class citizens, meaning you can pass them as arguments, assign them to variables, and return them from other functions.

This is powerful for scenarios where you want to define behavior dynamically, such as passing a function to another function for custom behavior.


*/
func add(a, b int) int {
	return a + b
}

func sub(a, b  int) int {
	return a - b
}

func main() {
	var calc func(int, int) int

	calc = add
	fmt.Println("Addition:", calc(10, 5))

	calc = sub
	fmt.Println("Subtraction:", calc(10, 5))

}
