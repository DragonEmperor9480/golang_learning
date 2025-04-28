package main

import "fmt"

//anonymous functions

func main() {
	func(name string) {
		fmt.Println("Hello", name)
	}("Linus")
}
