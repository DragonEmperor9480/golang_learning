/*
	Scope of Variables
	Package level: Declared outside any function. (Available to entire package.)
	Function level: Inside a function (Local).
	Block level: Inside {} braces.
*/

package main

import "fmt"

var name string = "One Punch Man" // can be accessed anywhere in the program

func main() {
	var name2 string = "Saitama" //can be accessed inside main function

	{
		var name3 string = "Genos"
		fmt.Println(name3) //can't be accessed outside block
	}
	fmt.Println(name,"\n", name2)
}
