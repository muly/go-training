package main

import "fmt"

func main() {
	//5accessingIndividualElement/2readingOutsideBoundary.go
	// accessing individual element: outside the boundary
	a := [4]int{1, 2, 3, 4}
	fmt.Println(a[4]) //ERROR: invalid argument: array index 4 out of bounds [0:4]
}
