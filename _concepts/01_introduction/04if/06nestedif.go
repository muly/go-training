package main

import "fmt"

func main() {
	var a, b int = 10, 2
	var class1, class2 = "go class", "java class"

	if a == b {
		if class1 == class2 {
			fmt.Println("all are equal")
		} else {
			fmt.Println("numbers are equal")
		}
	} else {
		fmt.Println("nothing is matching")
	}

}
