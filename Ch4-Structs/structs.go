/*
Structs are composite types(like classes in other languages but without inheritance)

they group fields together under a single type.

Why Structs?
- they help organize data
- can create custom types
- used heavily in APIs, structs are backbone of Go applications
- can have methods associated with them(like classes, but no inheritance)
*/

package main

import "fmt"

// Struct initilization
type Car struct {
	brand string
	year  int
}

type Cat struct {
	brand string
	name  string
	age   int
}

func main() {

	//struct
	var c1 Car
	c1.brand = "BMW"
	c1.year = 2025

	fmt.Println(c1.brand, c1.year)

	//Using Positional Initialization
	c2 := Cat{"Local", "Meoww", 5}
	fmt.Println(c2.brand, c2.name, c2.age)

	//Using Named Initialization (recommended)
	c3 := Cat{
		brand: "Local",
		name:  "Lab Sample",
		age:   1,
	}
	fmt.Println(c3)

}
