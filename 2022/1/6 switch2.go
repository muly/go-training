package main

import "fmt"

/*
switch variable{
case value1:
	// add your code
	fallthrough // optional
case value2:
	// add your code
default: // optional
	// add your code
}
*/

func main() {
	gender := "F" 
	/*
	switch {
	case gender == "M":
		fmt.Println("male")
	case gender == "F":
		fmt.Println("female")
	default:
		fmt.Println("any")
	}*/
	switch gender{
	case "M":
		fmt.Println("male")
	case "F":
		fmt.Println("female")
	default:
		fmt.Println("any")
	}

}
