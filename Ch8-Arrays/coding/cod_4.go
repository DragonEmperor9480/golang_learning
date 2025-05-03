/*
🧩 4. Check if Array is Palindrome
Task:
Write a function isPalindrome(arr [5]int) bool that checks if the array reads the same forward and backward.
*/

package main

import "fmt"

func isPalindrome(arr [5]int) bool {
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	num := [5]int{1, 2, 3, 2, 1}
	fmt.Println(isPalindrome(num))
}
