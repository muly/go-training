package main

import "fmt"

func main() {
	var a, b int = 10, 20

	if a > b {
		fmt.Println("a is larger")
	} else if b > a {
		fmt.Println("b is larger")
	} else {
		fmt.Println("numbers are same")
	}
}
