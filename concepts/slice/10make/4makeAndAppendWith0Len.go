package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 0, 4)

	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)

	a = append(a, 50)
	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)
}
