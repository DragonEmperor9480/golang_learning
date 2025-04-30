package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + "says woof!"
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

func main() {
	var a Animal

	d := Dog{
		Name: "Buddy",
	}
	a = d
	fmt.Println(a.Speak())

	c := Cat{
		Name: "PussPuss",
	}

	a = c
	fmt.Println(a.Speak())
}
