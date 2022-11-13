package main

import "fmt"

func main() {
	// assigning one slice to another slice
	s1 := []int{1, 2, 3, 4}
	var s2 []int
	s2 = s1
	fmt.Println(s1, s2)
}
