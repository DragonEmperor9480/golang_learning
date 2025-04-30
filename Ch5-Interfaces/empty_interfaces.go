package main

import "fmt"

func PrintAny(i interface{}) {
	fmt.Println(i)
}

func main() {
	PrintAny(42)
	PrintAny("Hello, Go")
	PrintAny(3.14)
}
