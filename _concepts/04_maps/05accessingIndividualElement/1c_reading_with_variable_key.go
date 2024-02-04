package main

import "fmt"

func main() {

	var m = map[int]string{1: "one", 3: "three", 0: "zero"}

	//
	//
	// START OMIT
	k := 1
	val := m[k] // same as m[1]
	fmt.Println(val)
	// END OMIT
}
