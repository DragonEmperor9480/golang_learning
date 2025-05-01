package main

import "fmt"

/*
Go has only one looping construct:
the for loop. It can be used in various ways to achieve all loop types (like while, do-while, and traditional for loops from other languages).

*/

func main() {

	fmt.Println("Basic For Loop")
	//Basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//While-Style Loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	//Infinite Loop
	for {
		fmt.Println("Inside an infinite loop")
		break
	}

	//Break and Continue
	for k := 1; k <= 5; k++ {
		if k == 3 {
			continue
		}

		if k == 5 {
			fmt.Println("5 found. Breaking....")
			break
		}
	}
}
