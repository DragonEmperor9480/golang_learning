package main

import "fmt"

func main() {
	// var m map[string]int   //decleration only(Nil map)
	// m = make(map[string]int) //Initialization

	m := map[string]int{
		"apples":  10,
		"bananas": 23,
	}

	fmt.Println(m)
	fmt.Println(m["apples"])
	fmt.Println(m["bananas"])

	m["coconuts"] = 12 //adding another element to the maps

	//Printing the values of the map
	fmt.Println(m)

	//accessing the values of map
	fmt.Println(m["apples"])
	fmt.Println(m["bananas"])
	fmt.Println(m["coconuts"])

	//Deleting the content of the map
	delete(m, "apples")

	fmt.Println(m)
	fmt.Println(m["apples"])
	fmt.Println(m["bananas"])
	fmt.Println(m["coconuts"])

	//Check whether the following key exists or not(it exists)
	val, ok := m["bananas"]
	fmt.Println(val)
	fmt.Println(ok)
	//(23, true)

	//check whether it exists or not
	val, ok = m["apples"]
	fmt.Println(val)
	fmt.Println(ok)
	//(0, false)

	//Length of the map
	fmt.Println(len(m))
}
