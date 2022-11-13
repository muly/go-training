package main

import "fmt"

func main() {
	// range: ignoring the value; simplified form
	s := []int{9, 8, 7, 6}
	for i := range s {
		fmt.Println(s[i])
	}
}
