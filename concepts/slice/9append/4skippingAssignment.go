package main

import "fmt"

func main() {
	//append: skipping assignment
	var a []int = []int{1, 2, 3, 4}
	fmt.Printf("len=%d  %v\n", len(a), a)

	append(a, 5) // ERROR: append(a, 5) (value of type []int) is not used
	fmt.Printf("len=%d  %v\n", len(a), a)
}
