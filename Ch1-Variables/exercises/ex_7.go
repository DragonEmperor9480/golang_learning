//Create a block scope variable inside {} and print it inside the block.

package main

import "fmt"

func main(){
	{
		var num int = 100
		fmt.Println(num)
	}
}