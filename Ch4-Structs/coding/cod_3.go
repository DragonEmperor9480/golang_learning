/*
Create a Person struct with fields name, age, and address.
Write a method ShowDetails that prints the details of the person.
In the main function, create a Person object and call ShowDetails to display the details.
*/

package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) ShowDetails() {
	fmt.Println("Name: ", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("Address:", p.address)
}

func main() {
	p1 := Person{
		name:    "John",
		age:     22,
		address: "Some Highway Road,Tokyo,Japan",
	}
	p1.ShowDetails()
}
