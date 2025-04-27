/*
Coding Q4:
Use fallthrough to show how it works when matching a case.
*/

package main

import "fmt"

func main() {
	num := 1
	switch num {
	case 1:
		fmt.Println("1 is the first number")
		fallthrough
	case 2:
		fmt.Println("2 comes after 1")
	default:
		fmt.Println("Error")
	}
}
