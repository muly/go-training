package main

import "fmt"

func main() {
	// for loop: length is hard coded
	s := []int{1, 2, 3, 4}
	for i := 0; i < 4; i++ {
		fmt.Println(s[i])
	}
}
