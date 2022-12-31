package main

import "fmt"

func main() {
	// reinitializing the slice
	s := []int{1, 2, 3, 4}

	s = []int{10, 30, 50, 60}
	fmt.Println(s)
}
