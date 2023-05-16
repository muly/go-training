package main

import "fmt"

var y = 10 // package level scope

func main() {

	var z = 10 // function level scope
	{
		var x int = 10 // block level scope
		fmt.Println("inside inner block",x, y, z)
		{

		}
	}
	{
		fmt.Println("inside inner block",x, y, z) // x is not visible here
	}
	
	fmt.Println("in main block", x, y, z) // x is not visible here
}

