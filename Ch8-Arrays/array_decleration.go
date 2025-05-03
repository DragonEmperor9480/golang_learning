package main

import "fmt"

//What is an Array?
// Array is a fixed size collection of elements of same type

func main() {
	var nums [3]int //an array of 3 integers [0 0 0] or if it was string then ""
	fmt.Println(nums)

	//Declaring and Initializing Arrays
	var a [3]int                                                 //zero-initialized
	b := [3]int{1, 2, 3}                                         //with values
	c := [...]string{"c", "c++", "Rust", "Go", "Java", "Python"} //inferred size
	fmt.Println("Zero-initialized:", a)
	fmt.Println("With Values:", b)
	fmt.Println("Inferred Size:", c)

	//Accessing and modifying elements
	b[1] = 100
	fmt.Println("Modified 'b' array value at index 1:", b[1])
	fmt.Println("Total array of b:", b)

	//Length of array
	fmt.Println("Length of Array B:", len(b))

	//Arrays are value types
	//When passed to functions or assigned to another variable, arrays are copied.
	x := [2]int{1, 2}
	y := x
	y[0] = 100

	fmt.Println(x) // [1 2] orginal unchanged
	fmt.Println(y) // [100 2]

	//Iteration over Arrays
	fmt.Println("Iteration of Arrays")
	for i, v := range c {
		fmt.Println(i+1,":", v)
	}
}
