package main

import "fmt"

func main() {

	var age int
	var ignoreTodler bool
	var ignoreBoy bool

	age = 10
	/*
		&& AND
		|| OR
		! NOT
	*/
	//
	if (age < 4) && (ignoreTodler) == false {
		fmt.Println("todler")
	} else if age < 13 && !ignoreBoy {
		fmt.Println("boy")
	} else if (age < 20) {
		fmt.Println("teen")
	} else {
		fmt.Println("adult")
	}

}

/*
	if boolVar == true{}
	// same as
	if boolVar{}
*/

/*
if boolVar == false{}
	// same as
if !boolVar{}
*/
