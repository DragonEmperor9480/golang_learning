/*
Write a function countWords(sentence string) map[string]int that returns a map with the frequency of each word in the given sentence.
*/

package main

import (
	"fmt"
	"strings"
)

func countWords(sentence string) map[string]int {
	wordMap := make(map[string]int)

	words := strings.Fields(sentence)

	for _, word := range words {
		wordMap[word]++
	}

	return wordMap

}

func main() {
	result := countWords("Go go go")
	fmt.Println(result)
}
