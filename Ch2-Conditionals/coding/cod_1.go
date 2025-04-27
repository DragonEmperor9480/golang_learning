/*
Coding Q1:
Write a program that checks if a number is positive, negative, or zero using if-else.
*/

package main

func main() {
	num := -5
	if num > 0 {
		print(num, "is positive")
	} else if num < 0 {
		println(num, "is negative")
	} else {
		println(num, "is zero")
	}

}
