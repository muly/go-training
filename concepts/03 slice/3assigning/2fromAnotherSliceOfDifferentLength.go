package main

import "fmt"

func main() {
	// assigning one slice to another array of different length.
	// Note: slice DON'T have the length as part of the data type
	s1 := []int{1, 2, 3, 4}
	var s2 []int = []int{3, 4, 5, 6, 7, 8}
	s2 = s1
	fmt.Println(s1, s2)
	

	// we can see more clear example of this concept
	// after we explore the concept of make() and len().
}
