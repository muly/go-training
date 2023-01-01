package main

import "fmt"

func main() {
	// 3assigning/2fromAnotherArrayOfDifferentLength.go
	// assigning one array to another array of different length. array has the length as part of the data type
	a := [4]int{1, 2, 3, 4}
	var b [5]int
	b = a // ERROR:
	fmt.Println(a, b)
}
