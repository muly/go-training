package main

import "fmt"

func main() {
	var a []int
	a = make([]int, 0, 4)

	fmt.Println(len(a), cap(a))
	fmt.Println(a)
}
