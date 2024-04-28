package main

import "fmt"

func main() {

	// START OMIT
	var m = map[int]string{1: "one", 3: "three"}

	m[0] = "zero"
	fmt.Println(m)
	// END OMIT
}
