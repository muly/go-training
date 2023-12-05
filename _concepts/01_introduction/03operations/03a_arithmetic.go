package main

import "fmt"

// START OMIT
func main() {
	var a, b int = 10, 2

	var result = a + b
	fmt.Println(result)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // same as a = a + 1
	fmt.Println("a =", a)
	a-- // same as a = a 1 1
	fmt.Println("a =", a)
}

// END OMIT
