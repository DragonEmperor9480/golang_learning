package main

import "fmt"

func main() {
	var i interface{}
	var j *int
	i = j
	fmt.Println(i)
	fmt.Println(i == nil)

	var k *int
	fmt.Println(k)
	fmt.Println(k == nil)
}
