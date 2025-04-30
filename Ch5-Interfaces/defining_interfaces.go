package main

import "fmt"

type Speaker interface {
	speak() string
}

type Person struct {
	name string
}

func (p Person) speak() string {
	return "Hello," + p.name
}

func main() {
	var speaker Speaker
	p := Person{
		name: "John",
	}

	speaker = p
	fmt.Println(speaker.speak())
}
