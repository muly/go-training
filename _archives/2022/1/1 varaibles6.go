package main

import "fmt"

// varaible shadowing?

func main() {
	var x = 99 // function level scope
	{
		var x int = 10 // block level scope
		fmt.Println("inside inner block",x)
	}
	
	fmt.Println("in main block", x)
}

