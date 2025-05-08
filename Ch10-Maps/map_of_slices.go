package main

import "fmt"

func main() {
	ms := map[string][]int{
		"prime":     {2, 5, 7},
		"not_prime": {4, 6, 8},
	}

	fmt.Println(ms["prime"])    // print the prime
	fmt.Println(ms["prime"][0]) // print the specific value of the  slice
}
