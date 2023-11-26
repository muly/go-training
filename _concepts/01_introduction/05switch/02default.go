package main

import "fmt"

func main() {
	var a, b int = 10, 20

	switch {
	case a == b:
		fmt.Println("numbers are same")
	default:
		fmt.Println("numbers are different")
	}
}
