package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	var b [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Printing array b")
	fmt.Println(b)

	fmt.Println("Length of b:", len(b))

	fmt.Println("updating the array")
	b[0], b[1] = 8, 9

	fmt.Println("Printing updating array")
	fmt.Println(b)

	fmt.Println("Looping through array")

	for i, e := range b {
		fmt.Println("Index:", i, " Element:", e)
	}

	//SLICES
	fmt.Println("SLICES")

	slice := make([]int, 5, 10)
	fmt.Println("slice:", slice)
	fmt.Println("len:", len(slice), " Capacity:", cap(slice))
	slice[0] = 0
	slice[1] = 1
	slice[2] = 2
	slice[3] = 3
	slice[4] = 4
	fmt.Println("Slice Updated:", slice)
	fmt.Println("len:", len(slice), " Capacity:", cap(slice))

	slice2 := make([]int, 5, 10)

	copied := copy(slice2, slice)
	fmt.Println("slice2:", slice2)
	fmt.Println(copied)

	fmt.Println("Appending to slice")
	slice = append(slice, 5)
	slice = append(slice, 6)
	slice = append(slice, 7)
	slice = append(slice, 8)
	slice = append(slice, 9)
	slice = append(slice, 5)

	fmt.Println("len:", len(slice), " Capacity:", cap(slice))

	fmt.Println("Appending Slices")
	slice = append(slice, slice2...)
	fmt.Println(slice)
	fmt.Println("len:", len(slice), " Capacity:", cap(slice))

	//merging two slices in one slice
	slice = append(slice, slice2...)
	fmt.Println("len:", len(slice), " Capacity:", cap(slice))

	//deleting a value from slice
	fmt.Println("slice2:", slice2)
	i := 2

	slice_1 := slice2[:i]
	fmt.Println(slice_1)
	slice_2 := slice2[i+1:]
	fmt.Println(slice_2)
	final_slice := append(slice_1, slice_2...)
	fmt.Println("Final slice", final_slice)

}
