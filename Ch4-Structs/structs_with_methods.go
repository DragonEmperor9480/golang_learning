package main

import "fmt"

type Car struct {
	brand string
	year  int
}

func (c Car) showDetails() {
	fmt.Println("Brand:", c.brand)
	fmt.Println("Year:", c.year)
}

func main() {
	car1 := Car{"BMW", 2025}
	car1.showDetails()

}


