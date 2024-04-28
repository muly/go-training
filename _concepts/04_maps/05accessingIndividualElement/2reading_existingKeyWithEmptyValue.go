package main

import "fmt"

func main() {

	// START OMIT

	var m = map[int]float32{1: 111, 3: 333, 9: 0}

	fmt.Println(m[9])
	fmt.Println(m)

	// END OMIT
	fmt.Println(m)
}
