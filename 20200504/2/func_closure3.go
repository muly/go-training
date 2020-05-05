package main

import (
	"fmt"
)

func main() {
	var a string = "Hello"
	message := func(b string) {
		a = a + " " + b
	}

	message("class")
	fmt.Println(a)

	message("world")
	fmt.Println(a)
}
