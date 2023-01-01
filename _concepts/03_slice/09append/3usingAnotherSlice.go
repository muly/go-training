package main

import "fmt"

func main() {
	// append: using another slice contents (... expansion notation)
	var a []int = []int{1, 2, 3, 4}
	fmt.Printf("len=%d  %v\n", len(a), a)

	var b = []int{5, 6, 7}

	a = append(a, b...)
	fmt.Printf("len=%d  %v\n", len(a), a)
}
