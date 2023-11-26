package main

import "fmt"

func main() {
	var a, b int = 10, 20

	switch {
	case a > b:
		fmt.Println("a is larger")
	case b > a:
		fmt.Println("b is larger")
	default:
		fmt.Println("numbers are same")
	}
}
