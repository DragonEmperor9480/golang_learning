package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

var game string = "Where Winds Meet"

func main() {

	var a int = 43
	fmt.Println(a)

	var b float32 = 54.65
	fmt.Println(b)

	var c string = "Dragon"
	fmt.Println(c)
	fmt.Println(unsafe.Sizeof(c)) // can use this to print size of datatype

	var d bool = false
	fmt.Println(d)
	fmt.Println(unsafe.Sizeof(d))

	fmt.Printf("%T\n", d) // %T Can be used to print the Type of datatype

	var user string
	user = "Raiden"
	user = "Shogun"

	fmt.Println(user)

	var name1, name2 string = "Yae", "Miko"

	fmt.Printf("%s %s\n", name1, name2)

	//Printing Global variable
	fmt.Println(game)

	//taking user input
	// fmt.Println("Enter a game name:")
	// fmt.Scanf("%s", &game)
	// fmt.Println(game)

	//multiple input
	// var (
	// 	name string
	// 	age  int
	// )
	// fmt.Println("Enter name and age:")
	// fmt.Scanf("%s%d", &name, &age)
	// fmt.Println(name, age)

	//scanf eror handling

	// fmt.Println("Enter Gennshin impact rating")
	// var rate int
	// count, err := fmt.Scanf("%d", &rate)
	// //here above count refers how many input it has taken suessfully incase of multiple inputs are taken in same scanf
	// fmt.Println(count, rate)
	// fmt.Println(reflect.TypeOf(count)) // tells what kind of datatype
	// if err != nil {
	// 	fmt.Println("MF you entered something wrong")

	// }

	//TYPE CASTINGGGGGGGG

	var num int = 10
	var numf float32
	numf = float32(num)
	fmt.Printf("%.7f\n", numf)

	var numf2 float64 = 123.123456
	num = int(numf2)
	fmt.Println(num)

	NewNum := strconv.Itoa(num)
	fmt.Println("string:", NewNum)

	fmt.Println("Converting str back to int")
	i, err := strconv.Atoi(NewNum)
	fmt.Println("output of str -> int:", i, err)

	//atoi iota
}
