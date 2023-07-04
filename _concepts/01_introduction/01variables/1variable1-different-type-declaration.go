package main

import "fmt"

// START OMIT
func main() {
	// 1. declare & initialize in separate lines
	var name string
	name = "go class"

	// 2. declare & initialize in same lines
	var class string = "go class"

	// declare with datatype inferred
	var anotherClass = "go class"

	// declare with shorthand notation
	yetAnotherClass := "go class"

	fmt.Println(name, class, anotherClass, yetAnotherClass)
}

// END OMIT
