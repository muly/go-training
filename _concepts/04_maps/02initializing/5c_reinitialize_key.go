package main

import "fmt"

func main() {
	//
	var m map[int]string
	m = map[int]string{} // empty initialization

	// initializing individual key value pairs
	//
	m[1] = "one"
	m[2] = "two"

	// START OMIT
	// re-initialize a key with another value
	m[1] = "ONEEEE" // update existing key with new value
	
	// later-initializing a new key value pair
	m[3] = "three"
	// END OMIT
	fmt.Println(m)
}
