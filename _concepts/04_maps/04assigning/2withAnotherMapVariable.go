package main

import "fmt"

func main() {

	// START OMIT
	var m map[int]string

	n := map[int]string{1: "one", 3: "three", 0: "zero"}

	m = n

	fmt.Println(m)
	// TODO: print the address of both maps
	// END OMIT
	fmt.Println(m)
}
