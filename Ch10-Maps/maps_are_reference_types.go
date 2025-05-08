package main

import "fmt"

func modify(m map[string]int) {
	m["new"] = 42
}

func main() {
	m := map[string]int{"x": 1}
	modify(m)
	fmt.Println(m) // map[x:1 new:42]
}
