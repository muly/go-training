package main

import "fmt"

func main() {
	// accessing individual element: outside the boundary
	s := []int{1, 2, 3, 4}
	fmt.Println(s[4]) //ERROR:
}
