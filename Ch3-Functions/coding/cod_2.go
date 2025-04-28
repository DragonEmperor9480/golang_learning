/*
Create a variadic function called sumAll that adds any number of integers passed and returns the sum.
*/

package main

import "fmt"

func sumAll(numbers ...int) int{
	total:=0
	for _,num:= range numbers{
		total+=num
	}
	return total
}

func main(){
	res:=sumAll(1,2,3,4,5,6,7,8,9,0)
	fmt.Println(res)
}