package main

import "fmt"

// START OMIT
var p = 100 // package level scope

func main() {

	var f = 10 // function level scope

	{ // this is the starting of the block
		var b int = 1 // block level scope
		{ // this is an inner block
			var b2 int =2
			fmt.Println("inside inner block 1:", p, f, b, b2)
		}

		fmt.Println("inside block 1:", p, f, b)
	}

	{ // this is another block
		fmt.Println("inside block 2:", p, f) // b is not visible in this block
	}

	fmt.Println("in main function:", p, f) // b is not visible here too
}

// END OMIT
