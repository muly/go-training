package main

import "fmt"

func main() {

	// START OMIT
	var m = map[int]string{1: "one", 3: "three", 0: "zero"}

	m[3] = "A" // key is unique. if not, overwrites existing value
	fmt.Println(m)

	m[1] = "A" // value need not be unique. key 3, key 1 both has the same value "A"
	fmt.Println(m)
	// END OMIT
}
