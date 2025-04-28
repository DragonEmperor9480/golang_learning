/*

Write a function that uses defer to print "Exiting function" after printing "Inside function".

*/

package main

import "fmt"

func deferEx(){
	defer fmt.Println("Exiting function")
	fmt.Println("Inside function")
	
}

func main(){
	deferEx()
}