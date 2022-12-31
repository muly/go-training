package main

import "fmt"

func main() {
	// append: multiple elements
	var a []int = []int{1, 2, 3, 4}
	fmt.Printf("len=%d  %v\n", len(a), a)

	a = append(a, 5, 6, 7)
	fmt.Printf("len=%d  %v\n", len(a), a)
}
