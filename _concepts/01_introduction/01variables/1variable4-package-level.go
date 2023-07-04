package main

import "fmt"

// START OMIT
// declare at package level and initialize later
var name string

// name = "go class" // NOT allowed: lines should start with Go keywords only

// declare at package level and initialize here
var class string = "go class"

// declare at package level with data type inferred and initialize here
var anotherClass = "go class"

// declare with shorthand notation NOT supported at package level
// yetAnotherClass := "go class" // syntax error: non-declaration statement outside function body

func main() {
	name = "go class"
	yetAnotherClass := "go class"

	fmt.Println(name, class, anotherClass, yetAnotherClass)
	anotherClass = "go class 2"

	fmt.Println("hello")
}

// END OMIT
