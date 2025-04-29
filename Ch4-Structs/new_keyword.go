package main

import "fmt"

type Car struct {
	brand string
	model string
	year  int
}

func main() {
	c1 := Car{
		brand: "Audi",
		model: "Unknown",
		year:  2021,
	}
	fmt.Println(c1)
	fmt.Println("\n")

	//using new keyword
	c2 := new(Car)
	c2.brand = "Maruti Suzuki"
	c2.model = "Alto 800"
	c2.year = 2010
	fmt.Println(c2)

	//Zero values in structs
	var c3 Car
	fmt.Println(c3)

}
