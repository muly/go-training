package main

import "fmt"

func main() {
	// range: ignoring the value
	s := []int{9, 8, 7, 6}
	for i, _ := range s {
		fmt.Println(s[i])
	}
}
