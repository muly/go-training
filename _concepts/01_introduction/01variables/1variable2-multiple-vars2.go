package main

import "fmt"

// START OMIT
func main() {
	// declaring multiple variables together
	var a, b int
	a = 10
	b = 12
	
	// initializing multiple variables together
	a, b = 10, 12
	fmt.Println(a, b)

	// swap
	a, b = b, a 
	fmt.Println(a, b)
}

// END OMIT
