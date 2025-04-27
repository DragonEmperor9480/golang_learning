/*
Coding Q5:
Write a program to determine if a character is a vowel or consonant.

*/

package main

import "fmt"

func main() {
	vowel := 'e'
	switch vowel{
	case 'a', 'e','i','o','u':
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}
}
