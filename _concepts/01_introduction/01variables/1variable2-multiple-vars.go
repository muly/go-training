// multiple variables together
package main

import "fmt"

func main() {
	var x int
	var y int
	x = 10
	y = 12
	
	// declaring multiple variables together
	var a, b int
	a = 10
	b = 12
	
	// initializing multiple variables together
	a, b = 10, 12

	fmt.Println(x, y, a, b)
}

// https://go.dev/tour/basics/11
