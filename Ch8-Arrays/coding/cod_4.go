/*
🧩 4. Check if Array is Palindrome
Task:
Write a function isPalindrome(arr [5]int) bool that checks if the array reads the same forward and backward.
*/

package main

import "fmt"

func isPalindrome(arr [5]int) bool {
	z := 0
	for i, _ := range arr {
		if arr[i] != arr[(len(arr)-1)-i] {
			if arr[(len(arr)-1)-i] <= 0 {
				break
			}
			z++
		}

	}
	if z > 0 {
		return false
	} else {
		return true
	}

}

func main() {
	num := [5]int{1, 2, 3, 2, 1}
	fmt.Println(isPalindrome(num))
}
