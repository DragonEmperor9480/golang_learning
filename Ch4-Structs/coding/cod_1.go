/*

Write a program to create a struct Rectangle with fields length and width.
Add a method Area that calculates the area of the rectangle and prints it.
Use this struct in the main function.

*/

package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func (a Rectangle) Area() {
	area := a.length * a.width
	fmt.Println("Area of Rectangle:", area)
}

func main() {
	a1 := Rectangle{10, 20}
	a1.Area()

}
