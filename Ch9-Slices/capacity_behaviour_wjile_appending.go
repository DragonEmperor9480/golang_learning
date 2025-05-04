package main

import "fmt"

func main() {
	s := make([]int, 2, 3)
	fmt.Println("Before:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, 4) //fits in old array
	fmt.Println("Affter 1 append:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, 5) //triggers reallocation
	fmt.Println("Affter 2 append:", s, "len:", len(s), "cap:", cap(s))

}
