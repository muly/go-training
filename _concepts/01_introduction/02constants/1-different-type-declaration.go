package main

import "fmt"

// START OMIT
func main() {
	// // declare & initialize in separate lines: NOT ALLOWED
	// const pi float32 // ERROR: missing init expr for pi
	// pi = 3.141 // ERROR: cannot assign to pi (constant unknown of type float32)
	// // NOTE: this is like re-initializing, which is not allowed for constants

	// 1. declare & initialize in same lines
	const class string = "go class"
	// class = "java class" // ERROR: cannot re-assign to class

	// 2. declare with datatype inferred
	const anotherClass = "go class"

	//declare with shorthand notation: NOT for constant
	yetAnotherClass := "go class" // this becomes a variable, NOT constant

	fmt.Println(class, anotherClass, yetAnotherClass)
}

// END OMIT
