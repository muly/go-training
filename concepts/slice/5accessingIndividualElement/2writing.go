package main

import "fmt"

func main() {
	// changing individual element:
	s := []int{1, 2, 3, 4}
	s[1] = 100
	fmt.Println(s)
}
