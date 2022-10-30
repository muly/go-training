package main

import "fmt"

func main() {

	var age int
	var ignoreBoy bool

	age = 10

	//
	if age < 4 {
		fmt.Println("todler")
		if ignoreBoy {
			fmt.Println("something else")
		}
	} else if age < 13 {
		fmt.Println("boy")
	} else if age < 20 {
		fmt.Println("teen")
	} else {
		fmt.Println("adult")
	}

}
