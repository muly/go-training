package main

import "fmt"

func main() {
	// range: ignoring the index
	s := []int{9, 8, 7, 6}
	for _, v := range s {
		fmt.Println(v)
	}
}
