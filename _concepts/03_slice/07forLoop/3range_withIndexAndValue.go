package main

import "fmt"

func main() {
	// range: intro
	s := []int{9, 8, 7, 6}
	for i, v := range s {
		fmt.Println(i, v)
	}
}
