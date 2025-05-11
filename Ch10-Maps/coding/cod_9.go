/*
9. Map of Slices
Create a map where each key is a student's name and each value is a slice of their grades. Write a function to calculate the average grade for a given student.
*/

package main

import "fmt"

func avgGrades(x map[string][]int) map[string]float64 {
	avgMap := make(map[string]float64)

	for student, grades := range x {
		sum := 0
		for _, grade := range grades {
			sum += grade
		}
		if len(grades) > 0 {
			avgMap[student] = float64(sum) / float64(len(grades))
		}
	}
	return avgMap

}

func main() {
	studentGrades := map[string][]int{
		"Alice":   {90, 85, 92},
		"Bob":     {78, 80, 79},
		"Charlie": {100, 95, 98},
	}

	averages := avgGrades(studentGrades)
	fmt.Println("Average Grades:", averages)
}
