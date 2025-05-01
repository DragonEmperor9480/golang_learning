/*

🧩 5. Panic and Recover Example
Task:
Write a function crashApp() that panics.
Use defer and recover in main() to catch and print the panic.

*/

package main

import "fmt"

func crashApp() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error!")
		}
	}()
	panic("An Error Occured")
}

func main() {
	crashApp()
}
