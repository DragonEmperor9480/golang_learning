/*
🧩 5. Count Even and Odd Numbers
Task:
Write a function countEvenOdd(arr [10]int) (int, int) that returns two integers: the number of even and odd elements in the array.
*/

package main

import "fmt"

func countEvenOdd(arr [10]int) (int, int) {
	even := 0
	odd := 0
	for _, x := range arr {
		if x == 0 {
			continue
		} else if x%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return even, odd

}

func main() {
	num := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	x, y := countEvenOdd(num)
	fmt.Printf("Even:%d\n", x)
	fmt.Printf("Odd:%d\n", y)
}
