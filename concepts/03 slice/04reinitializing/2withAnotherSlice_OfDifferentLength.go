package main

import "fmt"

func main() {
	// reinitializing the slice: with different length
	var s []int = []int{1, 2, 3, 4}

	s = []int{20, 30, 40, 50, 60, 70}
	fmt.Println(s)

	// NOTE: for a slice, the size is not part of the datatype. 
	// so slices of different size are compatible 
}
