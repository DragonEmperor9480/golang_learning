/*
Using a single composite literal, create a Person with embedded Address initialized.

Example:
type Address struct { city string }
type Person struct {
    name string
    Address
}

*/

package main

import "fmt"

type Address struct{ city string }

type Person struct {
	name string
	Address
}

func main() {
	p1 := Person{
		name: "Alhaitham",
		Address: Address{
			city: "Bengaluru",
		},
	}

	fmt.Println("Name", p1.name)
	fmt.Println("City:", p1.city)
}
