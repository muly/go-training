package main

import "fmt"

func main() {
	// for loop: using len()
	s := []int{1, 2, 3, 4}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
