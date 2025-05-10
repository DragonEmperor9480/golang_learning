/*

Merge Two Maps
Write a function mergeMaps(m1, m2 map[string]int) map[string]int that merges two maps. If a key exists in both, sum their values.

*/

package main

import "fmt"

func mergeMaps(m1, m2 map[string]int) map[string]int {

	mergedMap := make(map[string]int)

	for x, y := range m1 {
		mergedMap[x] = y

	}

	for x, y := range m2 {
		mergedMap[x] += y

	}

	return mergedMap
}

func main() {
	org1 := map[string]int{
		"IN": 1,
		"US": 2,
		"FR": 3,
	}

	org2 := map[string]int{
		"CN": 4,
		"UK": 5,
		"FR": 7,
	}

	fmt.Println(mergeMaps(org1, org2))
}
