package main

import "fmt"
//multiple Returns
func areaAndPerimeter(length int, breadth int) (int, int) {
	area := length * breadth
	perimeter := 2 * (length + breadth)
	return area, perimeter
}

func main() {
	x, y := areaAndPerimeter(12, 2)
	fmt.Printf("Area: %d\n", x)
	fmt.Printf("Perimeter: %d\n", y)

}
