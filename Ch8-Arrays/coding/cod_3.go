/*
🧩 3. Reverse an Array
Task:
Write a function reverseArray(arr [4]int) [4]int that returns a new array with elements reversed.
*/

package main

import "fmt"

func reverseArray(arr [4]int) [4]int {
	new_arr := [4]int{}
	for i, x := range arr {
		new_arr[(len(arr)-1)-i] = x
	}
	return new_arr
}

func main() {
	num := [4]int{1, 2, 3, 4}
	fmt.Println("Reversed Array:", reverseArray(num))
}
