package main

import "fmt"

func main() {
	fruits := map[int]string{
		1: "Apple",
		2: "Banana",
		3: "Mango",
		4: "Strawberry",
	}

	for key, val := range fruits {
		fmt.Println("Key:", key, "  ", "Value:", val)
	}
}
