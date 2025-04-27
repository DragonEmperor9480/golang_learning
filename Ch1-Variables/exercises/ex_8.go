// Try to use := outside a function and observe the error (important!)

package main

import "fmt"

city:= "Hyderabad"

func main(){
	fmt.Println(city)
}
// syntax error: non-declaration statement outside function body