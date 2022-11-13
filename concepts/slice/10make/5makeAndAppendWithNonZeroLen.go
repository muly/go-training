package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 2, 4)

	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)

	a = append(a, 50)
	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)

	a[1] = 60 // you can always change the content of the initial cells
	fmt.Printf("len=%d  cap=%d %v\n", len(a), cap(a), a)
}
