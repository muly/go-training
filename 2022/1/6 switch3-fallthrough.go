package main

import "fmt"

/*
switch {
case condition1:
	// add your code
	fallthrough // optional
case conditionn:
	// add your code
default: // optional
	// add your code
}
*/

func main() {
	age := 2
	switch {
	case age < 1:
		fmt.Println("baby")
	case age < 3:
		fmt.Println("todler")
		fallthrough
	case age < 13:
		fmt.Println("boy")
		fallthrough
	case age < 20:
		fmt.Println("teen")
	default:
		fmt.Println("adult")
	}
	//

}
