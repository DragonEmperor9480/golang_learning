package main

import (
	"fmt"
	"strings"
)

func main() {

	//Count frequency of integers in a slice.
	nums := []int{1, 2, 2, 3, 1, 4, 2}

	freq := make(map[int]int)

	for _, val := range nums {
		freq[val]++
	}
	fmt.Println(freq)

	//Count frequency of strings.

	langs := []string{"go", "java", "go", "python", "java"}

	str_freq := make(map[string]int)

	for _, x := range langs {
		str_freq[x]++
	}

	fmt.Println(str_freq)

	//Count character frequency in a string.

	s := "gopher"

	str_ctr := make(map[rune]int)

	for _, x := range s {
		str_ctr[x]++
	}

	fmt.Println(str_ctr)

	//Find the most frequent element in a slice.
	nums2 := []int{1, 3, 2, 3, 4, 3, 2}

	fre_ctr := make(map[int]int)

	for _, x := range nums2 {
		fre_ctr[x]++
	}
	fmt.Println(fre_ctr)
	var max int = 0
	var mostFrequent int = 0
	for key, val := range fre_ctr {
		if max < val {
			max = val
			mostFrequent = key

		}
	}

	fmt.Println(mostFrequent)

	//Find all elements that appear more than once.

	for x, val := range fre_ctr {
		if val > 1 {
			fmt.Println(x)
		}
	}

	//Count word frequency in a sentence
	sentence := "go is fast and go is simple"
	str := strings.Split(sentence, " ")
	// fmt.Println(str)

	str_count := make(map[string]int)
	for _, val := range str {
		str_count[val]++
	}

	fmt.Println(str_count)

	//Count HTTP status codes from logs
	logs := []int{200, 500, 404, 200, 500, 200, 401}

	log_freq := make(map[int]int)

	for _, val := range logs {
		log_freq[val]++
	}

	fmt.Println(log_freq)

	//Count duplicate usernames
	users := []string{"alice", "bob", "alice", "eve", "bob"}

	user_count := make(map[string]int)
	for _, val := range users {
		user_count[val]++
	}

	for key, val := range user_count {
		if val > 1 {
			fmt.Println(key)
		}
	}

	//First non-repeating character in a string
	s = "swiss"
	non_repeat := make(map[rune]int)
	for _, val := range s {
		non_repeat[val]++
	}

	for _, val := range s {
		if non_repeat[val] == 1 {
			fmt.Printf("%c\n", val)
			break

		}
	}

	for key, val := range non_repeat {
		if val == 1 {
			fmt.Printf("%c ", key)
			break
		}
	}

	//Find all repeating characters
	s = "programming"
	str_dup_count := make(map[rune]int)
	for _, val := range s {
		str_dup_count[val]++
	}

	fmt.Println(str_dup_count)

	for key, val := range str_dup_count {
		if val > 1 {
			fmt.Printf("%c ", key)
		}
	}

}
