package main

import "fmt"

type MyError struct{
	message string
	code int
}

func(e MyError) Error() string{
	return fmt.Sprintf("Code %d: %s",e.code,e.message)
}

func riskyOperation() error{
	return MyError{message: "Something went wrong",code:500}
}

func main(){
	err:= riskyOperation()
	if err != nil{
		fmt.Println("Custom Error:",err)
	}
}