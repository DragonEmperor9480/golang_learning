package main

import "fmt"

func main() {

	//explict type decleration
	var num1 int
	num1 = 5
	fmt.Println(num1)

	//var + initialization
	var age int = 25
	fmt.Println(age)

	//Type Interface
	var city = "Bengaluru"
	fmt.Printf("%s\n", city)

	//Short-hand := Operator(only inside functions i.e not inside functions)
	city2 := "Hubli"
	fmt.Println(city2)

	//multiple variables at once
	var a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	//constants(fixed values)be changed after being set
	//these values can't 
	const pi = 3.141

	/*
	Type | Default Zero Value
	int | 0
	float64 | 0.0
	string | "" (empty string)
	bool | false
	pointer | nil
	*/

	var x int
	fmt.Println(x) //output:0

}
