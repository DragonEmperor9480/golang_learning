/*
type Contact struct {
    phone string
}

type User struct {
    Contact
    phone string
}

In main, assign different values to both phone fields and print them to show how field shadowing works.

*/

package main

import "fmt"

type Contact struct {
	phone string
}

type User struct {
	Contact
	phone string
}

func main() {
	c1 := Contact{
		phone: "1234567890",
	}
	fmt.Println(c1.phone)

	c2 := User{
		Contact: Contact{
			phone: "0987654321",
		},
		phone: "9090909090",
	}

	fmt.Println(c2.phone)

}
