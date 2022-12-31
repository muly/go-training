package main

import "fmt"

func main() {
	// capacity of s slice is the maximum number of elements the slice can hold.
	// unlike the array, slice is flexible and can grow.

	// if we want to add more elements beyond the capacity of the slice, new slice needs to be created. 
	// i.e. new memory location will be allocated, and a new underlying array will be created.
	// append() function will do this automatically for us.
	//
	// in the below example, the length and capacity is same because of the way we are initializing it.
	// after we explore make() system function, we will see that we can create a slice whose size is less than the capacity
	var s []int
	s = []int{1, 2, 3, 4}
	fmt.Printf("len=%d  cap=%d %v\n", len(s), cap(s), s)
}
