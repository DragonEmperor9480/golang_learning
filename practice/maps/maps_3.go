package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Group numbers into EVEN and ODD

	nums := []int{1, 2, 3, 4, 5, 6}

	even_odd_map := make(map[string][]int)

	for _, val := range nums {
		if val%2 == 0 {
			even_odd_map["even"] = append(even_odd_map["even"], val)
		} else {
			even_odd_map["odd"] = append(even_odd_map["odd"], val)
		}
	}

	fmt.Println(even_odd_map)

	//Group words by first letter
	words := []string{"go", "java", "python", "golang", "js"}

	first_letter_map := make(map[rune][]string)
	for _, val := range words {
		first_letter := []rune(val)[0]
		first_letter_map[first_letter] = append(first_letter_map[first_letter], val)
	}

	for key, value := range first_letter_map {
		fmt.Printf("%c: %v\n", key, value)
	}

	//Group students by GRADE
	students := []struct {
		Name  string
		Grade string
	}{
		{"Amit", "A"},
		{"Ravi", "B"},
		{"Neha", "A"},
	}

	grade_map := make(map[string][]string)

	for _, val := range students {
		grade_map[val.Grade] = append(grade_map[val.Grade], val.Name)

	}

	fmt.Println(grade_map)

	len_map := make(map[int][]int)
	//Group numbers by DIGIT COUNT
	nums = []int{1, 22, 333, 44, 5}

	for _, val := range nums {
		x := strconv.Itoa(val)
		len_map[len(x)] = append(len_map[len(x)], val)
	}

	fmt.Println("Length of map")
	fmt.Println(len_map)

	//Bucket numbers into RANGES

	// 	Buckets

	// "0-10"

	// "11-20"

	// "21+"
	nums = []int{3, 8, 12, 17, 25, 30}

	bucket_map := make(map[string][]int)

	for _, val := range nums {
		if val >= 0 && val <= 10 {
			bucket_map["0-10"] = append(bucket_map["0-10"], val)

		} else if val >= 11 && val <= 20 {
			bucket_map["11-20"] = append(bucket_map["11-20"], val)

		} else {
			bucket_map["21+"] = append(bucket_map["21+"], val)
		}
	}

	fmt.Println("Bucket Map")
	fmt.Println(bucket_map)

}
