package main

import "fmt"

func greet(name string) string {
	return "Hello," + name
}

func main() {
	res := greet("Rias")
	fmt.Println(res)
}
