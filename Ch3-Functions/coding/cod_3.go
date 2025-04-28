/*
Write a named return function that returns the square of an integer.

*/

package main

import "fmt"

func sqnum(a int)(sqval int){
	sqval=a*a
	return
}

func main(){
	fmt.Println(sqnum(5))
}