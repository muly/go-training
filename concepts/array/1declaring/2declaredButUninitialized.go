package main

import "fmt"

func main() {
	// 1declaring/2declaredButUninitialized.go
	
	// un initialized variable
	var i int
	fmt.Println(i)

	// un initialized array: initialized all elements with zero value
	var a [4]int
	fmt.Println(a) // [0 0 0 0]
}
