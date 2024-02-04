package main

import "fmt"

func main() {
	//
	var m map[int]string
	m = map[int]string{} // empty initialization
	
	// initializing individual key value pairs
	// START OMIT
	m[1] = "one"
	m[2] = "two"
	// END OMIT
	fmt.Println(m)
}
