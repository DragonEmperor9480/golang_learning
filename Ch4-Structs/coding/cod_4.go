/*

Write a program that creates a struct Circle with a field radius.
Add a method Circumference that calculates the circumference of the circle.
In the main function, create a Circle object and print the circumference.

*/

package main

import "fmt"

type Circle struct {
	radius int
}

func (c Circle) Circumference() float64 {
	return 2 * 3.14 * float64(c.radius)
}

func main() {
	c1 := Circle{5}
	res := c1.Circumference()
	fmt.Printf("Circumference: %.2f\n", res)
}
