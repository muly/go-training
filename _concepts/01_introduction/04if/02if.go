package main

import "fmt"

func main() {
	var a, b int = 10, 10
	var class1, class2 = "go class", "go class"

	if (a == b) && (class1 == class2) {
		fmt.Println("all are equal")
	}
}
