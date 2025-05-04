//copying slices

package main

import "fmt"

// n := copy(destSlice, srcSlice)
//where `n` is thr number of elements acutally copied

func main() {
	src := []int{1, 2, 3}
	dest := make([]int, 3)
	n := copy(dest, src)
	fmt.Println("copied", n, "elements:", dest)

	// if destination is smaller than source only it will copy the first elements until dest becomes full
	dest2 := make([]int, 2)
	n = copy(dest2, src)
	fmt.Println(dest2)

	//Overlapping slices
	new_slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	copy(new_slice[1:], new_slice[:4])
	fmt.Println(new_slice)

}

/*
copy() only works on slices, not arrays.

copy() is useful when:
	Avoiding mutation of original data
	Pre-allocating memory for performance
*/