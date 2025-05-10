/*
4. Remove Key From Map
Write a function removeKey(m map[string]int, key string) that removes a given key from the map if it exists.
*/
package main

import "fmt"

func removeKey(m map[string]int, key string) {
	delete(m, key)
}

func main() {
	org1 := map[string]int{
		"IN": 1,
		"US": 2,
		"FR": 3,
		"CN": 4,
		"UK": 5,
	}

	fmt.Println("Before function:", org1)

	removeKey(org1, "CN")
	fmt.Println("After function:", org1)

}
