/*
Coding Q3:
Write a program to check if someone is:

Under 18: "Minor"

18–59: "Adult"

60 and above: "Senior Citizen"
*/

package main

import "fmt"

func main() {
	age := 30

	if age < 18 {
		fmt.Println("Minor")
	} else if age < 60 {
		fmt.Println("Adult")
	} else {
		fmt.Println("Senior Citizen")
	}
}
