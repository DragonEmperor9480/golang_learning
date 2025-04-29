/*
Create a struct Employee with fields id, name, and salary.
Write a method RaiseSalary that increases the salary by a given percentage.
In the main function, create an Employee object and print the original and updated salary.
*/

package main

import "fmt"

type Employee struct {
	id     int
	name   string
	salary float64
}

func (e Employee) RaiseSalary(perc float64) float64 {
	return e.salary + (e.salary * (perc / 100))
}

func main() {
	e1 := Employee{1568, "John", 1000.00}

	res := e1.RaiseSalary(10)
	fmt.Println("Orginal Salary:", e1.salary)
	fmt.Println("Updated Salary:", res)
}
