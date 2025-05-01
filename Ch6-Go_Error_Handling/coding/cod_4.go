/*
🧩 4. User Login Error
Task:
Create a Login function that accepts username and password.
If username != "admin" or password != "1234", return a proper error using fmt.Errorf.
*/

package main

import "fmt"

func Login(username, password string) error {
	if username != "admin" || password != "1234" {
		return fmt.Errorf("The username %s and password %s is wrong", username, password)
	} else {
		return nil
	}
}

func main() {
	res := Login("admin", "12345")
	if res != nil {
		fmt.Println(res)
	} else {
		fmt.Println("Login sucessful")
	}
}
