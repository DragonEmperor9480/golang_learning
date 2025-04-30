/*
panic() is used when something goes very wrong — like accessing invalid memory.

recover() lets you recover from a panic, preventing the program from crashing.
*/

package main

import "fmt"

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}

	}()
	panic("Critical Error occured!")

}

func main() {
	riskyFunction()
	fmt.Println("Program continues...")
}
