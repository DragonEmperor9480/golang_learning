/*

Create a Book struct with fields title, author, and price.
Write a method DiscountPrice that takes a percentage and calculates the discounted price.
In the main function, create a Book object and print the original and discounted price.

*/

package main

import "fmt"

type Book struct {
	title  string
	author string
	price  float64
}

func (p Book) DiscountPrice(percentage float64) float64 {
	return p.price - (p.price * (percentage / 100))

}

func main() {

	b1 := Book{
		title:  "Advanced Java",
		author: "Amrutesh Naregal",
		price:  1000,
	}

	discount_price := b1.DiscountPrice(10)
	fmt.Println("Orginal Price:", b1.price)
	fmt.Println("Discounted Price:", discount_price)

}
