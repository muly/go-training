package main

import "fmt"

func main() {
	var a, b int = 10, 2

	if a > b {
		fmt.Println("a is larger")
	} else if a < b {
		fmt.Println("b is larger")
	} else {
		fmt.Println("both are equal")
	}
}
