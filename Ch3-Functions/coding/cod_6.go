/*

Create an anonymous function that multiplies two numbers and call it inside main.

*/

package main

import "fmt"

func main() {
	func(a, b int) {
		fmt.Println(a * b)
	}(10, 20)
}
