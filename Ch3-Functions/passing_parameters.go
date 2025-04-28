package main

import "fmt"

//Functions without return

func addNumbers(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func main() {
	addNumbers(10, 10)
}
