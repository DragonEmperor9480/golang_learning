/*
Coding Q2:
Create a program using switch to print the month name for a given month number (1-12).

*/

package main

import "fmt"

func main(){
	month :=7
	switch month{
	case 1:
		fmt.Println("January")
	case 2:
		fmt.Println("Frburary")
	case 3:
		fmt.Println("March")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("May")
	case 6:
		fmt.Println("June")
	case 7:
		fmt.Println("July")
	case 8:
		fmt.Println("August")
	case 9:
		fmt.Println("Sepetember")
	case 10:
		fmt.Println("October")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("December")
	default:
		fmt.Println("Error!")
	}
}