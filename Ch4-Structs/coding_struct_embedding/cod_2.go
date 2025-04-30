/*

🧩 2. Method with Embedded Struct
Add a method FullAddress() to Address that prints both city and zipcode.

In main, create a Person and call FullAddress() through the Person object.

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

func (p Address) FullAddress() {
	fmt.Println("City:", p.city)
	fmt.Println("Zipcode:", p.zipcode)

}

func main() {
	p1 := Person{
		Address: Address{
			city:    "Mordor",
			zipcode: 100001,
		},
		name: "Shadow",
	}
	p1.FullAddress()

}
