package main

import (
	"errors"
	"fmt"
)

/*

🧩 3. File Read Simulation
Task:
Write a function readFile(fileName string) error.
If the filename is not "data.txt", return a FileNotFoundError (your custom type).
Otherwise, print "Reading file...".

*/

func readFile(filename string) error {
	if filename == "data.txt" {
		return nil
	} else {
		return errors.New("FileNotFoundError")
	}
}

func main() {
	res := readFile("data.txt")
	if res != nil {
		fmt.Println(res)
	} else {
		fmt.Println("Reading file...")
	}
}
