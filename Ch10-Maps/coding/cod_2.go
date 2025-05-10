/*

 2. Invert a Map
Write a function invertMap(m map[string]string) map[string]string that inverts a map (i.e., keys become values and values become keys).

*/

package main

import "fmt"

func invertMap(m map[string]string) map[string]string {
	revMap := make(map[string]string)

	for key, value := range m {
		revMap[value] = key
	}
	return revMap
}

func main() {
	org := map[string]string{
		"IN": "India",
		"US": "United Stats",
		"FR": "France",
	}

	fmt.Println(invertMap(org))
}
