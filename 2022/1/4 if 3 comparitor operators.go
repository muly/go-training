package main

import "fmt"

func main() {

	var age int

	age = 10

	// comparitor operators
	// <, >, ==, <=, >=, !=



	if age < 4 {
		fmt.Println("todler")
	} 
	if age < 13 {
		fmt.Println("boy")
	} 
	if age < 20 {
		fmt.Println("teen")
	} else {
		fmt.Println("adult")
	}


	if age < 4 {
		fmt.Println("todler")
	} else if age < 13 {
		fmt.Println("boy")
	} else if age < 20 {
		fmt.Println("teen")
	} else {
		fmt.Println("adult")
	}

	

}
