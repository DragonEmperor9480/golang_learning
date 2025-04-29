package main

import "fmt"

type Person struct {
	name   string
	salary float64
}

type Employee struct {
	Person
	emp_id int
}

func (e Employee) showDetails() {
	fmt.Println(e.name)
	fmt.Println(e.salary)
	fmt.Println(e.emp_id)
}

func main() {

	e2 := Employee{
		Person: Person{
			name:   "Amrut",
			salary: 5000.00,
		},
		emp_id: 1,
	}
	e2.showDetails()
}
