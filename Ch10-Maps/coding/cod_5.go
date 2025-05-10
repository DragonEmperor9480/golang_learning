/*
 5. Check Key Existence
Write a function keyExists(m map[int]string, key int) bool that returns true if the key exists in the map.

*/

package main

import "fmt"

func keyExists(m map[int]string, key int) bool {
	_, hehe := m[key]

	return hehe
}

func main() {
	numbers := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
	}
	fmt.Println(keyExists(numbers, 4))
	fmt.Println(keyExists(numbers, 5))

}
