/*
8. Top Scorer
Given a map of names to scores, write a function topScorer(scores map[string]int) string that returns the name of the student with the highest score.
*/

package main

import "fmt"

func topScorer(scores map[string]int) string {
	x, y := "", 0
	for key, val := range scores {
		if y < val {
			y = val
			x = key
		}
	}
	return x
}

func main() {
	map2 := map[string]int{
		"Arno":   98,
		"Akshay": 100,
		"Akash":  78,
	}
	fmt.Println(topScorer(map2))

}
