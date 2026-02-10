package main

import "fmt"

/*

Create a map[string]int to store product prices and print it.

Insert 5 key–value pairs into a map and update one existing key.

Read a value from a map where the key may not exist.

Use value, ok := map[key] to check key existence.

Delete a key from a map and print the map.

Print the length of a map before and after insertions.

Iterate over a map and print keys and values.

Demonstrate reading from a nil map.

Show what happens when you write to a nil map (commented explanation).

Compare an empty map vs a nil map using len.

*/

func main() {

	//Create a map[string]int to store product prices and print it.

	prod_price := make(map[string]int)

	fmt.Println(prod_price)
	fmt.Println(len(prod_price))

	//Insert 5 key–value pairs into a map and update one existing key.

	prod_price["apple"] = 79
	prod_price["mango"] = 90
	prod_price["strawberry"] = 67
	prod_price["orange"] = 34
	prod_price["cherry"] = 32
	prod_price["apple"] = 119

	fmt.Println(prod_price)

	//Read a value from a map where the key may not exist.
	//Use value, ok := map[key] to check key existence.

	val, ok := prod_price["banana"]

	if ok {
		print(val)
	} else {
		fmt.Println("not found")
	}

	//Delete a key from a map and print the map.

	delete(prod_price, "mango")
	fmt.Println(prod_price)

	//Print the length of a map before and after insertions.
	fmt.Println(len(prod_price)) //already printed len in beginning

	// Iterate over a map and print keys and values.

	for key, value := range prod_price {
		fmt.Println(key, "->", value)
	}

	//Demonstrate reading from a nil map.

	var new_map map[string]int
	fmt.Println(new_map)

	// Show what happens when you write to a nil map (commented explanation).
	//new_map["id"] = 56
	// Well it compiles succesfully but the code panics wen ran

	// Compare an empty map vs a nil map using len.
	//x := make(map[string]int)
	//fmt.Println(x==new_map)

}
