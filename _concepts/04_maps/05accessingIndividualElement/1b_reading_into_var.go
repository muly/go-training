package main

import "fmt"

func main() {

	var m = map[int]string{1: "one", 3: "three", 0: "zero"}

	// START OMIT
	val := m[1] // hardcoded key
	fmt.Println(val) // saving the value retrieved into a variable
	// END OMIT
}
