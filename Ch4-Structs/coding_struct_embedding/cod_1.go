/*
🧩 1. Embedded Struct Field Access
Create two structs:

# Address with fields city and zipcode

Person which embeds Address and has a field name
In main, create a Person object and print all the details using direct field access.
*/
package main

import "fmt"

type Address struct {
	city    string
	zipcode int
}

type Person struct {
	Address
	name string
}

func main() {
	p1 := Person{
		Address: Address{
			city:    "Mordor",
			zipcode: 100001,
		},
		name: "Shadow",
	}

	fmt.Println("Name:", p1.name)
	fmt.Println("City:", p1.city)
	fmt.Println("Zipcode:", p1.zipcode)
}
