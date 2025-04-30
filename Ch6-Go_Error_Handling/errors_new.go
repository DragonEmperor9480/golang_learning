package main

import (
	"errors"
	"fmt"
)

func getFile(filename string) error {
	if filename == "" {
		return errors.New("Filename cannot be empty")
	}
	return nil
}

func main() {
	err := getFile("")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
