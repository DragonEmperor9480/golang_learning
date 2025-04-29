package main

import "fmt"

// Anonymous functions are useful for temporary one-time data

func main() {
	student := struct {
		name string
		age  int
	}{
		name: "John",
		age:  20,
	}
	fmt.Println(student)
}
