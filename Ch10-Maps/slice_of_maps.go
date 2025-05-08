package main

import (
	"fmt"
)

func main() {
	sm := []map[string]string{
		{"hello": "world", "city": "NA"},
		{"hello": "Babylon", "city": "persia"},
	}
	fmt.Println(sm)

	sm = append(sm, map[string]string{"hello": "Amrut", "city": "Bengaluru"})
	fmt.Println(sm)
	sm = append(sm, map[string]string{"hello": "Sam", "city": "Earth"})
	fmt.Println(sm)
}

