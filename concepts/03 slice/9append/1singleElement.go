package main

import "fmt"

func main() {
	// append: a single element
	var s []int = []int{1, 2, 3, 4}
	fmt.Printf("len=%d  %v\n", len(s), s)

	s = append(s, 5)
	fmt.Printf("len=%d  %v\n", len(s), s)
}
