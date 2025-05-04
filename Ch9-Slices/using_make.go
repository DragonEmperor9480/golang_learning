package main

import "fmt"

func main() {
	slice := make([]int, 5) //[0 0 0 0 0]
	fmt.Println(len(slice)) //5
	fmt.Println(cap(slice)) //5

	//Lets see diff b/w length and capacity

	slice2 := make([]int, 2, 5) // [0 0]
	fmt.Println(len(slice2))    //2
	fmt.Println(cap(slice2))    //5

	//Appending slices
	slice2 = append(slice2, 10, 20, 30)
	fmt.Println(slice2)

	//what if you exceed the capacity?
	slice3 := make([]int, 2, 5) // [0 0]
	slice3 = append(slice3, 99, 9, 98, 89, 77, 776, 67)
	fmt.Println(slice3)
	
	//When you exceed the capacity, Go creates a new underlying array (often doubling the size).
	//Old data is copied to the new array.

}
