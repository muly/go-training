package main

import "fmt"

// START OMIT
func main() {
	var x = 99 // function level scope
	{
		var x int = 10 // block level scope, with same variable
		fmt.Println("inside inner block", x)
	}

	fmt.Println("in main block", x) // has value declare at function level
}

// END OMIT
