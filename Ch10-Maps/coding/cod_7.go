/*
 7. Map Length Check
Write a function isMapEmpty(m map[string]int) bool that returns true if the map is empty.
*/

package main

import "fmt"

func isMapEmpty(m map[string]int) bool {
	if len(m) == 0 {
		return true
	}
	return false
}

func main() {
	map1 := map[string]int{}
	map2 := map[string]int{
		"Hello": 1,
		"Hi":    2,
	}
	fmt.Println(isMapEmpty(map1))
	fmt.Println(isMapEmpty(map2))
}
