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

	age := 20
	if age < 1 {
		fmt.Println("baby")
	} else if age < 3 {
		fmt.Println("todler")
	} else if age < 13 {
		fmt.Println("boy")
	} else if age < 20 {
		fmt.Println("teen")
	} else {
		fmt.Println("adult")
	}

	switch {
	case age < 1:
		fmt.Println("baby")
	case age < 3:
		fmt.Println("todler")
	case age < 13:
		fmt.Println("boy")
	case age < 20:
		fmt.Println("teen")
	default:
		fmt.Println("adult")
	}
	//

}
