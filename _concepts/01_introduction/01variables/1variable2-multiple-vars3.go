package main

import "fmt"

// declaring multiple variables in group

// START OMIT
func main() {
	var (
		a    int
		name string
	)

	a = 10
	name = "go class"

	fmt.Println(a, name)
}

// END OMIT
