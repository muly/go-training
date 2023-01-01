package main

import "fmt"

func main() {
	// 3assigning/1fromAnotherArray.go
	// assigning one array to another array
	a := [4]int{1, 2, 3, 4}
	var b [4]int
	b = a
	fmt.Println(a, b)
}
