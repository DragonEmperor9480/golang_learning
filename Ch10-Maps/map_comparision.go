package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m == nil) //true
	m = make(map[string]int)
	m["new"] = 42
	fmt.Println(m == nil) //false
}
