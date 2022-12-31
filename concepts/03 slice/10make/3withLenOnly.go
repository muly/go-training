package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 2)

	fmt.Println(len(a), cap(a))
	fmt.Println(a)
}
