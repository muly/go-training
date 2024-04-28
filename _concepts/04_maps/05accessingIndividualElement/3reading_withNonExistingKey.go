package main

import "fmt"

func main() {

	// START OMIT
	//
	var m = map[int]float32{1: 111, 3: 333}

	fmt.Println(m[9]) // no error. returns zero value for the value type, i.e. 0 for float32
	fmt.Println(m)
	//
	// END OMIT
	fmt.Println(m)
}
