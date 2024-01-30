package main

import "fmt"

// START OMIT
var p1 int = 10

func main() {
	fmt.Println("in main block, but package level scope", p1)
	var p1 int = 20
	myfunc()

	fmt.Println("in main block, function level scope", p1)
}

func myfunc() {
	fmt.Println("in package level scope", p1)
}

// END OMIT
