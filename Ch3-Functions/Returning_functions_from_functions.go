/*
Returning Functions from Functions

In Go, functions can also return other functions.
This feature is useful when you want to create a function dynamically based on parameters or other conditions.
*/

package main

import "fmt"

func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	multiplyBy3 := multiplyBy(3)
	fmt.Println(multiplyBy3(5)) //output will be 15

	multiplyBy10:=multiplyBy(10)
	fmt.Println(multiplyBy10(5))//output will be 50
}
