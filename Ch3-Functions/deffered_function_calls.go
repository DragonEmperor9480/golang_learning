package main

import "fmt"

/*
Deferred Function Calls
Go provides the defer keyword, which allows you to schedule a function call to be executed after the surrounding function completes. 
It's often used for cleanup operations.

*/
func main() {
	defer fmt.Println("This will be printed last")
	fmt.Println("This will be printed first")
	fmt.Println("this will be printed second")
}
