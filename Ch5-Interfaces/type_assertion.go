package main

import "fmt"

//A type assertion allows you to retrieve the underlying value of an interface and assert its concrete type.

func printType(i interface{}) {
	v, ok := i.(int)
	if ok {
		fmt.Println("Integer Value:", v)
	} else {
		fmt.Println("Not an integer")
	}

}

func main() {
	var x interface{} = 42
	printType(x)
	x = "Hello"
	printType(x)
}
