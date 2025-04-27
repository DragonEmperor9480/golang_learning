package main

import "fmt"

func main() {

	//If else example
	age := 18
	if age >= 18 {
		fmt.Println("You are eligible to vote!")
	} else {
		fmt.Println("You are not eligible to vote")
	}

	//else if ladder

	marks := 48

	if marks >= 90 {
		fmt.Println("A+")
	} else if marks >= 75 {
		fmt.Println("A")
	} else if marks >= 50 {
		fmt.Println("B+")
	} else {
		fmt.Println("B")
	}
}
