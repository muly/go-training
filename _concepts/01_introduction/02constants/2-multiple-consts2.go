package main

import "fmt"

// START OMIT
func main() {

	// declare and initialize multiple variables together
	const (
		class string = "go class"
		year  int = 2023
	)
	fmt.Println(class, year)
}

// END OMIT
